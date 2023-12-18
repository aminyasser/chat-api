package main

import (
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	apiRoutes := router.Group("/api/v1")
	{
		// Golang service routes
		apiRoutes.POST("/apps/:app_token/chats", golangServiceHandler)
		apiRoutes.POST("/apps/:app_token/chats/:chat_number/messages", golangServiceHandler)

		//Rails service routes
		apiRoutes.GET("/apps/:app_token", railsServiceHandler)
		apiRoutes.GET("/apps", railsServiceHandler)
		apiRoutes.POST("/apps", railsServiceHandler)
		apiRoutes.GET("/apps/:app_token/chats", railsServiceHandler)
		apiRoutes.POST("/apps/:app_token/chats/:chat_number/messages/search", railsServiceHandler)

	}

	router.Run(":8000")
}

// Forward request to Golang service
func golangServiceHandler(c *gin.Context) {
	proxyRequest(c, "http://golang-app:8080")
}

// Forward request to Rails service
func railsServiceHandler(c *gin.Context) {
	proxyRequest(c, "http://rails-app:3000")
}

func proxyRequest(c *gin.Context, target string) {
	url, _ := url.Parse(target)
	// Setup reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	originalPath := c.Request.URL.Path
	// Removing '/api/v1' from the path
	modifiedPath := strings.TrimPrefix(originalPath, "/api/v1")
	c.Request.URL.Path = modifiedPath

	proxy.ServeHTTP(c.Writer, c.Request)
}
