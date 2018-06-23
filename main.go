package main

import (
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

	router.GET("/handshake", func(c *gin.Context) {
		verifyToken := os.Getenv("VERIFY_TOKEN")
		challenge := c.Query("hub.challenge")
		mode := c.Query("hub.mode")
		facebookVerifyToken := c.Query("hub.verify_token")

		if mode == "subscribe" && facebookVerifyToken == verifyToken {
			c.String(200, challenge)
		} else {
			c.String(403, "Unnautorized")
		}
	})
	return router
}
