package usecase

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/aminyasser/chat-api/golang-service/app/repository"
	"github.com/aminyasser/chat-api/golang-service/clients/redis"
	"github.com/aminyasser/chat-api/golang-service/domain/dto"
	"github.com/aminyasser/chat-api/golang-service/queue/producer"

	"github.com/gin-gonic/gin"
)

var (
	ErrAppTokenDoesntExist   error = errors.New("app token doesn't exist")
	ErrCannotConnectToRedis  error = errors.New("cannot connect to redis")
	ErrCannotReadCredentials error = errors.New("cannot read credentials")
)

func CreateChat(ctx *gin.Context) {
	token := ctx.Param("app_token")

	// check if token exists in apps table
	if exists, _ := repository.AppTokenExists(token); !exists {
		ctx.Error(ErrAppTokenDoesntExist)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"status": false, "message": ErrAppTokenDoesntExist.Error()})
		return
	}

	redis := redis.GetRedis()
	chatNumber, err := redis.Incr(ctx, token).Result()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, ErrCannotConnectToRedis)
		return
	}

	chat, err := json.Marshal(dto.Chat{
		Number:   chatNumber,
		AppToken: token,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, ErrCannotReadCredentials)
		return
	}

	// rabbitmq logic here
	producer.Produce("chats_queue", chat)

	ctx.JSON(201, gin.H{
		"message":     "Chat created successfully",
		"chat_number": chatNumber,
	})
}

func UpdateAppChatCount() {

	err := repository.UpdateAppChatsCount()
	if err != nil {
		log.Print(err.Error())
		return
	}
	log.Print("Updating apps.chat_count finished.")

}

func UpdateChatMessagesCount() {

	err := repository.UpdateChatMessagesCount()
	if err != nil {
		log.Print(err.Error())

		return
	}

	log.Print("Updating chats.message_count finished.")

}
