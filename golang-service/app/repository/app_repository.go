package repository

import (
	db "github.com/aminyasser/chat-api/golang-service/clients/mysql"
	"github.com/aminyasser/chat-api/golang-service/domain/model"
)

func AppTokenExists(token string) (bool, error) {
	app := model.App{}
	conn, err := db.GetDB()
	if err != nil {
		return false , err
	}
	err = conn.First(&app, "app_token = ?", token).Error 
	if err != nil {
		return false, err
	}
	return true, nil
}


func UpdateAppChatsCount() (err error) {
	conn, err := db.GetDB()
	if err != nil {
		return err
	}
	err = conn.Exec("UPDATE apps SET chat_count = (SELECT COUNT(id) FROM chats WHERE app_token= apps.app_token)").Error
	return err
}
