package main

import (
	"github.com/alendavid/containers/services/bakery-app/pkg/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(http.CORSMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
