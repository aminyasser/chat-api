package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/aminyasser/chat-api/golang-service/app/usecase"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	serviceRoutes := r.Group("/api/v1")
	{
		serviceRoutes.POST("/apps/:app_token/chats" , usecase.CreateChat)
		// serviceRoutes.POST("/apps/:app_token/chats/:chat_number/messages")
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}
