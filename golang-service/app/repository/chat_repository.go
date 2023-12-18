package repository

import (
	db "github.com/aminyasser/chat-api/golang-service/clients/mysql"
	"github.com/aminyasser/chat-api/golang-service/domain/model"
)

func CreateChat(chat model.Chat)  error {
	conn, err := db.GetDB()
	if err != nil {
		return err
	}
	return conn.Create(&chat).Error
}
