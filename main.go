package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func init() {
	/*
	   Safety net for 'too many open files' issue on legacy code.
	   Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	   Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	*/
	http.DefaultClient.Timeout = time.Second * 10
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/alive", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "alive",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		temp := uuid.NewV4() //strings.Replace(c.Request.URL.Path, "/", "", 2)
		//log.Println("*******", c.Request.URL.Path)
		c.JSON(200, gin.H{
			"data": temp,
		})
	})

	r.Run(":8081")
}
