package service

import (
	"encoding/json"
	"errors"
	gdb "gin/repository/gorm"

	// repository "gin/repository/mysql"
	rdb "gin/repository/redis"

	"gin/model"

	"github.com/redis/go-redis/v9"
)

type bookService struct {
	bookRepo gdb.BookRepositoryGORM
	// bookRepo repository.BookRepository
	bookRedis rdb.BookRepositoryRedis
}

func NewBookService(bookRepo gdb.BookRepositoryGORM,
	bookRedis rdb.BookRepositoryRedis) bookService {
	return bookService{
		bookRepo:  bookRepo,
		bookRedis: bookRedis,
	}
}

func (s bookService) GetAllBook() ([]BookResponse, error) {
	books, err := s.bookRepo.GetAllBook()
	if err != nil {
		return nil, err
	}
	result := []BookResponse{}
	for _, book := range books {
		a := BookResponse{
			ID:   book.ID,
			Name: book.Name,
			Desc: book.Desc,
		}
		result = append(result, a)
	}

	return result, nil
}

func (s bookService) GetByID(id int) (*BookResponse, error) {
	var val BookResponse
	err := s.bookRedis.Get(id, &val)
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		book, err2 := s.bookRepo.GetByID(id)
		if err2 != nil {
			return nil, errors.New("book not found")
		}
		result := BookResponse{
			ID:   book.ID,
			Name: book.Name,
			Desc: book.Desc,
		}

		json, err3 := json.Marshal(result)
		if err3 != nil {
			return nil, err3
		}

		data := rdb.BookRedis{
			Key:        book.ID,
			Value:      json,
			Expiration: 0,
		}

		err4 := s.bookRedis.Set(data)
		if err4 != nil {
			return nil, err4
		}
		return &result, nil
	}
	return &val, nil
}

func (s bookService) GetBookAuthor(id int) (*AuthorResponse, error) {
	aut, err := s.bookRepo.GetBookAuthor(id)
	if err != nil {
		return nil, errors.New("failed to get book's author")
	}

	author := AuthorResponse{
		ID:   aut.ID,
		Name: aut.Name,
	}
	return &author, nil
}

func (s bookService) GetAuthorBook(id int) ([]BookResponse, error) {
	var result []BookResponse
	books, err := s.bookRepo.GetAuthorBook(id)
	if err != nil {
		return nil, errors.New("author not found")
	}

	for _, book := range books {
		a := BookResponse{
			ID:   book.ID,
			Name: book.Name,
			Desc: book.Desc,
		}
		result = append(result, a)
	}
	return result, nil
}

func (s bookService) UpdateBook(id int) (model.MessageResponse, error) {
	err := s.bookRedis.Delete(id)
	if err != nil {
		return model.MessageResponse{}, err
	}

	book, err2 := s.bookRepo.UpdateBook(id)
	if err2 != nil {
		return model.MessageResponse{}, errors.New("book not found")
	}

	// result := BookResponse{
	// 	ID:   book.ID,
	// 	Name: book.Name,
	// 	Desc: book.Desc,
	// }

	return book, nil

}

func (s bookService) AddBook() (model.MessageResponse, error) {
	book, err := s.bookRepo.AddBook()
	if err != nil {
		return model.MessageResponse{}, errors.New("failed to add new book")
	}

	// result := BookResponse{
	// 	ID:   book.ID,
	// 	Name: book.Name,
	// 	Desc: book.Desc,
	// }
	return book, nil
}

func (s bookService) DeleteBook(id int) (model.MessageResponse, error) {
	err := s.bookRedis.Delete(id)
	if err != nil {
		return model.MessageResponse{}, err
	}

	message, err := s.bookRepo.DeleteBook(id)
	if err != nil {
		return model.MessageResponse{}, errors.New("book not found")
	}
	return message, nil
}
