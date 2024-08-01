package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context) { // public harus angka besar
	c.JSON(http.StatusOK, gin.H{
		"content": "home",
	})
}

func (h *bookHandler)HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "hello content",
	})
}

func (h *bookHandler)BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func (h *bookHandler)QueryHandler(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func (h *bookHandler)PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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
	book, err := h.bookService.Create(bookRequest)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return 
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data" : book,
	})
}
