package repository

import (
	db "github.com/aminyasser/chat-api/golang-service/clients/mysql"
	"github.com/aminyasser/chat-api/golang-service/domain/model"
)

func CreateChat(chat model.Chat) (model.Chat, error) {
	conn, err := db.GetDB()
	if err != nil {
		return model.Chat{}, err
	}
	err = conn.Create(&chat).Error
	if err != nil {
		return model.Chat{}, err
	}
	return chat, nil
}
