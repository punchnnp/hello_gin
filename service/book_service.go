package service

import (
	"errors"
	repository "gin/repository/mysql"
	rdb "gin/repository/redis"
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
	book, err := s.bookRepo.GetByID(id)
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

func (s bookService) UpdateBook(id int) (*BookResponse, error) {
	book, err := s.bookRepo.UpdateBook(id)
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

func (s bookService) DeleteBook(id int) (string, error) {
	_, err := s.bookRepo.DeleteBook(id)
	if err != nil {
		return "", errors.New("book not found")
	}
	return "book deleted", nil
}
