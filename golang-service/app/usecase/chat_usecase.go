package usecase

import (
	"encoding/json"

	"github.com/aminyasser/chat-api/golang-service/clients/redis"
	"github.com/aminyasser/chat-api/golang-service/domain/dto"
	"github.com/aminyasser/chat-api/golang-service/queue/producer"

	"github.com/gin-gonic/gin"
)

// type MyController struct {
//     redisClient *redis.Client
// }

// func NewMyController(client *redis.Client) *MyController {
//     return &MyController{
//         redisClient: client,
//     }
// }
// func (c *MyController) SomeAction(ctx context.Context) {
//     // Use c.redisClient here
//     // ...
// }

func CreateChat(ctx *gin.Context) {
	token := ctx.Param("app_token")
	
	// check if token exists in apps table

	redis := redis.GetRedis()
	chatNumber, err := redis.Incr(ctx, token).Result()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Internal Server Error",
		})
	}

	chat, err := json.Marshal(dto.Chat{
		Number:   chatNumber,
		AppToken: token,
	})
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Internal Server Error",
		})
	}

	// rabbitmq logic here
	producer.Produce("chats_queue", chat)

	ctx.JSON(200, gin.H{
		"message":     "Chat created successfully",
		"chat_number": chatNumber,
	})
}
