package main

import (
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
	return router
}
