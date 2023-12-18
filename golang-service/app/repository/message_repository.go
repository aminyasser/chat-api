package repository

import (
	db "github.com/aminyasser/chat-api/golang-service/clients/mysql"
	"github.com/aminyasser/chat-api/golang-service/domain/model"
)

func CreateMessage(message model.Message)  error {
	conn, err := db.GetDB()
	if err != nil {
		return err
	}
	return conn.Create(&message).Error
}

func GetMessageChatId(token string , chatNumber int64)  (int , error) {
	chat := model.Chat{}
	conn, err := db.GetDB()
	if err != nil {
		return 0 ,err
	}
	res := conn.First(&chat, "app_token = ? AND chat_number = ?", token , chatNumber)
	if res.Error != nil {
		return 0 , res.Error
	}
	 
	return chat.Id , nil
}

