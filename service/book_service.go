package service

import (
	"encoding/json"
	"errors"
	repository "gin/repository/mysql"
	rdb "gin/repository/redis"
	"log"

	"gin/model"

	"github.com/redis/go-redis/v9"
)

type bookService struct {
	bookRepo  repository.BookRepository
	bookRedis rdb.BookRepositoryRedis
}

func NewBookService(bookRepo repository.BookRepository,
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
			log.Println("here")
			return nil, err4
		}
		return &result, nil
	}
	return &val, nil
}

func (s bookService) UpdateBook(id int) (*BookResponse, error) {
	err := s.bookRedis.Delete(id)
	if err != nil {
		return nil, err
	}

	book, err2 := s.bookRepo.UpdateBook(id)
	if err2 != nil {
		return nil, errors.New("book not found")
	}

	result := BookResponse{
		ID:   book.ID,
		Name: book.Name,
		Desc: book.Desc,
	}

	return &result, nil

}

func (s bookService) AddBook() (*BookResponse, error) {
	book, err := s.bookRepo.AddBook()
	if err != nil {
		return nil, errors.New("book not found")
	}

	result := BookResponse{
		ID:   book.ID,
		Name: book.Name,
		Desc: book.Desc,
	}
	return &result, nil
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
