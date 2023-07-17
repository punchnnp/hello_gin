package gorm

import (
	"errors"
	"gin/model"

	"gorm.io/gorm"
)

type bookRepositoryGORM struct {
	db *gorm.DB
}

func NewRepositoryGORM(db *gorm.DB) bookRepositoryGORM {
	return bookRepositoryGORM{db: db}
}

func (r bookRepositoryGORM) GetAllBook() ([]Books, error) {
	var books []Books
	result := r.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r bookRepositoryGORM) GetByID(id int) (*Books, error) {
	var book Books
	err := r.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r bookRepositoryGORM) GetBookAuthor(id int) (*Books, error) {
	var book Books
	result := r.db.Preload("Author").Find(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (r bookRepositoryGORM) GetAuthorBook(id int) ([]Books, error) {
	var books []Books
	result := r.db.Find(&books, "book_aut = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	row := result.RowsAffected
	if row == 0 {
		return nil, errors.New("author not found")
	}
	return books, nil
}

func (r bookRepositoryGORM) UpdateBook(id int) (model.MessageResponse, error) {
	new := Books{
		ID:   id,
		Name: "name updated",
		Desc: "description updated",
	}
	result := r.db.Where("book_id = ?", id).Updates(&new)
	if result.Error != nil {
		return model.MessageResponse{}, result.Error
	}
	// First, Last, Take have error return, else using RowsAffected
	row := result.RowsAffected
	if row == 0 {
		return model.MessageResponse{}, errors.New("book not found")
	}
	return model.MessageResponse{Message: "updated successfully"}, nil
}

func (r bookRepositoryGORM) AddBook() (model.MessageResponse, error) {
	new := Books{
		Name: "new book",
		Desc: "new description",
	}
	err := r.db.Create(&new).Error
	if err != nil {
		return model.MessageResponse{Message: "failed to add new book"}, err
	}
	return model.MessageResponse{Message: "add new book success!"}, nil
}

func (r bookRepositoryGORM) DeleteBook(id int) (model.MessageResponse, error) {
	var book Books
	result := r.db.Where("book_id = ?", id).Delete(&book)
	if result.Error != nil {
		return model.MessageResponse{}, result.Error
	}
	row := result.RowsAffected
	if row == 0 {
		return model.MessageResponse{}, errors.New("book bot found")
	}
	return model.MessageResponse{Message: "deleted success"}, nil
}
