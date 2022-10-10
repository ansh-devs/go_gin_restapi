package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// 	driver "anshdevs.in/golang/go/driver"
	// 	handler "anshdevs.in/golang/go/handler"
	// 	models "anshdevs.in/golang/go/models"
	// 	repository "anshdevs.in/golang/go/repository"
)

func main() {
	// driver.Message()
	// handler.Message()
	// models.Message()
	// repository.Message()
	r := gin.Default()
	r.TrustedProxies = nil
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{
				"message": "pong",
			})
	})
	r.StaticFile("/json", "./jsn.json")
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	var a string = "This is a string.."
	var b = &a
	println(b)
}
