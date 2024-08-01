package handler

import (
	"fmt"
	"strconv"

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
// find all datas
func (h *bookHandler) GetBooks(c *gin.Context) { // public harus angka besar
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : books,
	})
}
// find by id
func(h *bookHandler)GetBookByID(c *gin.Context){
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	
	book, err := h.bookService.FindByID(id)
	if err != nil {
		fmt.Println("eror")
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data" : book,
	})
	
}
// delete
func(h *bookHandler)DeleteBook(c *gin.Context){
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	
	book, err := h.bookService.FindByID(id)
	if err != nil {
		fmt.Println("eror")
		return
	}
	
	delete, err := h.bookService.Delete(book.ID)
	if err != nil{
		fmt.Print("error deleted book")
	}
	c.JSON(http.StatusOK, gin.H{
		"massage" : "delete book",
		"data" : delete,
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
