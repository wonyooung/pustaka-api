package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
	"fmt"
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
   // CRUD
   // book := book.Book{}
   // book.Title = "Atomic habits"
   // book.Price = 12000
   // book.Discount = 20
   // book.Rating = 5
   // book.Description = "Ini adalah buku yang sangan laris"
   
   // err = db.Create(&book).Error
   
   // if err != nil {
   // 	fmt.Println("Db connection error")
   // }
   	var book book.Book
   	err = db.Where("id = ?", 1).First(&book).Error
    if err != nil {
    	fmt.Println("finding book record")
    }
    fmt.Println(book.Title)
    book.Title = "Man tiger (revisi)"
    err = db.Save(&book).Error
    if err != nil {
    	fmt.Println("error updating book", book.ID)
    }
    

   
	router := gin.Default()
	
	v1 := router.Group("/v1")
	
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)
	
	router.Run(":111")
}


