package main

import (
	"encoding/json"
	"log"

	"github.com/aminyasser/chat-api/golang-service/app/repository"
	"github.com/aminyasser/chat-api/golang-service/clients/rabbitmq"
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
		logOnError(err, "Failed to register a consumer")
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
	<-forever
}
