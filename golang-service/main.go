package main

import (
	router "github.com/aminyasser/chat-api/golang-service/routes"
	"log"
)

func main() {
	router := router.InitRouter()
	router.Run()
	log.Println("Listening on 8080 ......")
}
