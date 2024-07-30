package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBooksHandler)

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

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}
func queryHandler(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required,number"`
}

func postBooksHandler(c *gin.Context) {
	// title, price
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
