package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const facebookAPI = "https://graph.facebook.com/v2.6/me/messages?access_token=%s"

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
				if messaging := entry.Messaging[0]; messaging.Message.Text != "" {
					log.Printf("Message: %v", messaging.Message.Text)
					sendMessage(messaging.Sender.ID, "42")
				}
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

func sendMessage(recipient string, text string) {
	messaging := Messaging{
		Message: Message{
			Text: text,
		},
		Recipient: User{
			ID: recipient,
		},
	}
	url := fmt.Sprintf(facebookAPI, os.Getenv("PAGE_ACCESS_TOKEN"))
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(messaging)
	res, err := http.Post(url, "application/json; charset=utf-8", body)

	if err != nil {
		log.Printf("Error to send message: %v", err)
	} else {
		log.Printf("Message sent: %v", res)
	}
}
