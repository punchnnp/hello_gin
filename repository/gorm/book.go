package gorm

import "gin/model"

type Books struct {
	ID   int    `gorm:"column:book_id" json:"id"`
	Name string `gorm:"column:book_name" json:"name"`
	Desc string `gorm:"column:book_desc" json:"description"`
}

type BookRepositoryGORM interface {
	GetAllBook() ([]Books, error)
	GetByID(int) (*Books, error)
	UpdateBook(int) (model.MessageResponse, error)
	AddBook() (model.MessageResponse, error)
	DeleteBook(int) (model.MessageResponse, error)
}
