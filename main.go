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
    bookRequest := book.BookRequest{
    	Title: "manusia harimau",
     	Price: 90909,
    }
    
    bookService.Create(bookRequest)
    
	router := gin.Default()
	
	v1 := router.Group("/v1")
	
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)
	
	router.Run(":01")
}


