package model

import "time"

type App struct {
	Id        int       `gorm:"column:id;primary_key;auto_increment" json:"-"`
	AppToken  string    `gorm:"column:app_token;not null" json:"app_token"`
	ChatCount int       `gorm:"column:chat_count;default:0" json:"chat_count"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"create_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}
