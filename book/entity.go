package book

import "time"

type Book struct{
	ID int
	Title string
	Description string
	Price int
	Rating int
	CreateAt time.Time
	UpdateAt time.Time
}