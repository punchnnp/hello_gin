package handler

import "gin/model"

// move interface to handler level cause
// handler is the one who define the details what service should be able to do
// then service implement according interface from handler
type bookHandler struct {
	bookService BookService
}

type BookResponse struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Desc   string `db:"description"`
	Author string `json:",omitempty"`
}

type AuthorResponse struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type BookService interface {
	GetAllBook() ([]BookResponse, error)
	GetByID(int) (*BookResponse, error)
	GetBookAuthor(int) (*BookResponse, error)
	GetAuthorBook(int) ([]BookResponse, error)
	UpdateBook(int) (model.MessageResponse, error)
	AddBook() (model.MessageResponse, error)
	DeleteBook(int) (model.MessageResponse, error)
}
