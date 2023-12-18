package main

import (
	"encoding/json"
	"log"

	"github.com/aminyasser/chat-api/golang-service/app/repository"
	"github.com/aminyasser/chat-api/golang-service/app/usecase"

	"github.com/aminyasser/chat-api/golang-service/clients/rabbitmq"
	"github.com/jasonlvhit/gocron"

	// "github.com/aminyasser/chat-api/golang-service/clients/redis"
	"github.com/aminyasser/chat-api/golang-service/domain/dto"
	"github.com/aminyasser/chat-api/golang-service/domain/model"
)

func logOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func main() {
	channel := rabbitmq.GetRabbitMQConsumeChannel()

	chats, err := channel.Consume(
		"chats_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logOnError(err, "Failed to register chat consumer")
	}

	messages, err := channel.Consume(
		"messages_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logOnError(err, "Failed to register message consumer")
	}

	forever := make(chan bool)

	go func() {
		for chat := range chats {
			var chatMessage model.Chat
			err := json.Unmarshal(chat.Body, &chatMessage)
			if err != nil {
				log.Print(err.Error())
				continue
			}

			err = repository.CreateChat(chatMessage)
			if err != nil {
				log.Print(err.Error())
				continue
			}

			if err == nil {
				log.Print("Chat Created Succefully")
			}
		}
	}()

	go func() {
		for message := range messages {
			var msgMessage dto.Message
			err := json.Unmarshal(message.Body, &msgMessage)
			if err != nil {
				log.Print(err.Error())
				continue
			}

			msg := model.Message{
				Number: int(msgMessage.Number),
				Body:   msgMessage.Body,
				ChatId: msgMessage.ChatId,
			}
			err = repository.CreateMessage(msg)
			if err != nil {
				log.Print(err.Error())
				continue
			}

			if err == nil {
				log.Print("Message Created Succefully")
			} else {
				
			}
		}
	}()
	go func() {
		gocron.Every(45).Minutes().Do(usecase.UpdateAppChatCount)
		gocron.Every(45).Minutes().Do(usecase.UpdateChatMessagesCount)
		<-gocron.Start()
	}()
	<-forever
}
