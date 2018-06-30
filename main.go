package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := ConfigureRouter()
	router.Run()
}

func ConfigureRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.GET("/webhook", handshake)
	router.POST("/webhook", processMessage)
	return router
}

func handshake(c *gin.Context) {
	verifyToken := os.Getenv("VERIFY_TOKEN")
	challenge := c.Query("hub.challenge")
	mode := c.Query("hub.mode")
	facebookVerifyToken := c.Query("hub.verify_token")

	if mode == "subscribe" && facebookVerifyToken == verifyToken {
		c.String(200, challenge)
	} else {
		c.String(403, "Unnautorized")
	}
}

func processMessage(c *gin.Context) {
	var event Event
	if err := c.ShouldBindJSON(&event); err == nil {
		if event.Object == "page" {
			for _, entry := range event.Entries {
				log.Printf("Entry received: %v", entry)
			}
		} else {
			log.Printf("Event not supported: %v", event.Object)
			c.String(404, "Event not supported")
		}
		c.String(200, "Message received")
	} else {
		log.Printf("Body: %v", c)
		log.Printf("Erro: %v", err)
		c.String(200, "Message received")
	}
}
