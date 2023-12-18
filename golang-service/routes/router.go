package router

import (
	"net/http"

	"github.com/aminyasser/chat-api/golang-service/app/usecase"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/apps/:app_token/chats", usecase.CreateChat)
	r.POST("/apps/:app_token/chats/:chat_number/messages", usecase.CreateMessage)

	// for api gatway wait-for-it script
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	return r
}
