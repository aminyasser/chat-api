package dto

type Message struct {
	RedisToken   string `json:"redis_token"`
	Number     int64  `json:"message_number"`
	Body       string `json:"body"`
	ChatId     int    `json:"id"`
}

type MessageBody struct {
	Body string `json:"body" binding:"required"`
}
