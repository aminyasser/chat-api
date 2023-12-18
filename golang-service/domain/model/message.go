package model

import "time"

type Message struct {
	Id        int       `gorm:"column:id;primary_key;auto_increment" json:"-"`
	ChatId    int       `gorm:"column:chat_id;not null" json:"-"`
	Number    int       `gorm:"column:message_number;not null" json:"chat_number"`
	Body      string    `gorm:"column:body;type:text;not null" json:"body"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"create_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}
