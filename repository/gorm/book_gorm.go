package gorm

import (
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
	// when query not existing id, it's not err and return {0}
	var book Books
	result := r.db.Find(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (r bookRepositoryGORM) UpdateBook(id int) (model.MessageResponse, error) {
	// update not existing id, not receive error
	new := Books{
		ID:   id,
		Name: "name updated",
		Desc: "description updated",
	}
	result := r.db.Where("book_id = ?", id).Save(&new)
	if result.Error != nil {
		return model.MessageResponse{}, result.Error
	}
	return model.MessageResponse{Message: "updated successfully"}, nil
}

func (r bookRepositoryGORM) AddBook() (model.MessageResponse, error) {
	// auto increment not work, book_id doesn't have default value
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
	// delete not existing id, not receive error
	var book Books
	result := r.db.Where("book_id = ?", id).Delete(&book)
	if result.Error != nil {
		return model.MessageResponse{}, result.Error
	}
	return model.MessageResponse{Message: "deleted success"}, nil
}
