package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.Run(":3000")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "home",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "hello content",
	})
}
