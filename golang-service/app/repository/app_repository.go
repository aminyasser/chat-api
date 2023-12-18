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
