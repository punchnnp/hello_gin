package mysql

import "gin/model"

type Book struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
}

type BookRepository interface {
	GetAllBook() ([]Book, error)
	GetByID(int) (*Book, error)
	UpdateBook(int) (*Book, error)
	AddBook() (*Book, error)
	DeleteBook(int) (model.MessageResponse, error)
}
