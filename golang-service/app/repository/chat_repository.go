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


func UpdateChatMessagesCount() (err error) {
	conn, err := db.GetDB()
	if err != nil {
		return err
	}
	err = conn.Exec("UPDATE chats SET message_count = (SELECT COUNT(id) FROM messages WHERE chat_id= chats.id)").Error
	return err
}
