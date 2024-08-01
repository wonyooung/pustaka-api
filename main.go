package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
   
	if err != nil {
   		log.Fatal("db connection error")
	}
  
   	db.AutoMigrate(&book.Book{}) // auto migrate 
  
    bookRepository := book.NewRepository(db)
    bookService := book.NewService(bookRepository)
    bookhandler := handler.NewBookHandler(bookService)
    
	router := gin.Default()
	
	v1 := router.Group("/v1")
	
	v1.GET("/", bookhandler.RootHandler)
	v1.GET("/hello", bookhandler.HelloHandler)
	v1.GET("/books/:id/:title", bookhandler.BooksHandler)
	v1.GET("/query", bookhandler.QueryHandler)
	v1.POST("/books", bookhandler.PostBooksHandler)
	
	router.Run(":01")
}


