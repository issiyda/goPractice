package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin")
	})
	route.Run(":8080")
}