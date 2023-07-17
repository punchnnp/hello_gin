package service

import "gin/model"

type BookResponse struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Desc string `db:"description"`
}

type AuthorResponse struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type BookService interface {
	GetAllBook() ([]BookResponse, error)
	GetByID(int) (*BookResponse, error)
	GetBookAuthor(int) (*AuthorResponse, error)
	GetAuthorBook(int) ([]BookResponse, error)
	UpdateBook(int) (model.MessageResponse, error)
	AddBook() (model.MessageResponse, error)
	DeleteBook(int) (model.MessageResponse, error)
}
