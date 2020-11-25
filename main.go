package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET(os.Getenv("loader"), func(c *gin.Context) {
		c.String(http.StatusOK, os.Getenv("loader"))
	})

	r.Run()
}
