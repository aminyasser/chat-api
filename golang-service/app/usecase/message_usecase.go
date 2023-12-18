package usecase

import (
	"encoding/json"
	"errors"
	"strconv"

	"net/http"

	"github.com/aminyasser/chat-api/golang-service/app/repository"
	"github.com/aminyasser/chat-api/golang-service/clients/redis"
	"github.com/aminyasser/chat-api/golang-service/domain/dto"
	"github.com/aminyasser/chat-api/golang-service/queue/producer"

	"github.com/gin-gonic/gin"
)

var (
	ErrMessageBodyDidNotExist error = errors.New("message body did not exist")
	ErrChatNumberDoesntExist  error = errors.New("chat number doesn't exist")
)

func CreateMessage(ctx *gin.Context) {
	var jsonData dto.MessageBody

	if err := ctx.ShouldBindJSON(&jsonData); err != nil {
		ctx.Error(ErrMessageBodyDidNotExist)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"status": false, "message": ErrMessageBodyDidNotExist.Error()})
		return
	}

	token := ctx.Param("app_token")

	if exists, _ := repository.AppTokenExists(token); !exists {
		ctx.Error(ErrAppTokenDoesntExist)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"status": false, "message": ErrAppTokenDoesntExist.Error()})
		return
	}

	chatNumber := ctx.Param("chat_number")
	number, _ := strconv.ParseInt(chatNumber, 10, 64)
	chatId, err := repository.GetMessageChatId(token, number)
	if err != nil {
		ctx.Error(ErrChatNumberDoesntExist)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"status": false, "message": ErrChatNumberDoesntExist.Error()})
		return
	}

	redis := redis.GetRedis()
	// TODO: decrement redis on error
	messageToken := chatNumber + ":" + token
	messageNumber, err := redis.Incr(ctx, messageToken).Result()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, ErrCannotConnectToRedis)
		return
	}

	message, err := json.Marshal(dto.Message{
		Number: messageNumber,
		Body:   jsonData.Body,
		ChatId: chatId,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, ErrCannotReadCredentials)
		return
	}

	producer.Produce("messages_queue", message)

	ctx.JSON(201, gin.H{
		"message":        "Message created successfully",
		"message_number": messageNumber,
	})
}
