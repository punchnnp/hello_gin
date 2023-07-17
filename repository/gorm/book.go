package gorm

import "gin/model"

type Books struct {
	ID     int     `gorm:"column:book_id;primary_key" json:"id"`
	Name   string  `gorm:"column:book_name" json:"name"`
	Desc   string  `gorm:"column:book_desc" json:"description"`
	Author *Author `gorm:"foreignKey:ID" json:",omitempty"`
	Aut    int     `gorm:"column:book_aut;"`
}

type Author struct {
	ID   int    `gorm:"column:aut_id"`
	Name string `gorm:"column:aut_name"`
}

func (author *Author) TableName() string {
	return "author"
}

type BookRepositoryGORM interface {
	GetAllBook() ([]Books, error)
	GetByID(int) (*Books, error)
	GetBookAuthor(int) (*Books, error)
	GetAuthorBook(int) ([]Books, error)
	UpdateBook(int) (model.MessageResponse, error)
	AddBook() (model.MessageResponse, error)
	DeleteBook(int) (model.MessageResponse, error)
}

//
