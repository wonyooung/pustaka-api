package book

import "time"

type Book struct{
	ID int
	Title string
	Description string
	Price int
	Rating int
	Discount int
	CreateAt time.Time
	UpdateAt time.Time
}