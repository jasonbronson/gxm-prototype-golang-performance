package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		temp := strings.Replace(c.Request.URL.Path, "/", "", 1)
		log.Println("*******", c.Request.URL.Path)
		c.String(http.StatusOK, temp)
	})

	r.Run()
}
