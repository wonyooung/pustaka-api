package book

type BookRequest struct {
	Title string `json:"title" binding:"required"`
	Price int `json:"price" binding:"required,number"`
}