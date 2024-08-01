package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) { // public harus angka besar
	c.JSON(http.StatusOK, gin.H{
		"content": "home",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "hello content",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func PostBooksHandler(c *gin.Context) {
	var bookInput book.BookRequest

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : errorMessages,
		})
		return
	}
	
	
	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
