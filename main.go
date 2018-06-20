package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := configureRouter()
	router.Run()
}

func configureRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/webhook", func(c *gin.Context) {
		verifyToken := os.Getenv("VERIFY_TOKEN")
		challenge := c.Params.ByName("hub.challenge")
		mode := c.Params.ByName("hub.mode")
		facebookVerifyToken := c.Params.ByName("hub.verify_token")

		if mode == "subscribe" && facebookVerifyToken == verifyToken {
			c.String(200, challenge)
		} else {
			c.String(403, "Unnautorized")
		}

	})
	return router
}
