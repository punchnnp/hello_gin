package repository

import (
	"errors"
)

type bookRepositoryMock struct {
	books []Book
}

func NewBookRepositoryMock() *bookRepositoryMock {
	books := []Book{
		{ID: 1, Name: "Book one", Desc: "This is book one"},
		{ID: 2, Name: "Book two", Desc: "This is book two"},
	}
	return &bookRepositoryMock{books: books}
}

func (m *bookRepositoryMock) GetAllBook() ([]Book, error) {
	if len(m.books) != 0 {
		return m.books, nil
	}
	return nil, errors.New("no book")
}

func (m *bookRepositoryMock) GetByID(id int) (*Book, error) {
	for _, book := range m.books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, errors.New("book not found")
}

func (m *bookRepositoryMock) UpdateBook(id int) (*Book, error) {
	for i := 0; i < len(m.books); i++ {
		if m.books[i].ID == id {
			m.books[i].Name = "Name change"
			m.books[i].Desc = "Description change"
			return &m.books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func (m *bookRepositoryMock) AddBook() (*Book, error) {
	newId := m.books[len(m.books)-1].ID + 1
	newBook := Book{
		ID:   newId,
		Name: "New book",
		Desc: "New description",
	}
	m.books = append(m.books, newBook)
	return &m.books[len(m.books)-1], nil
}

func (m *bookRepositoryMock) DeleteBook(id int) (string, error) {
	for i := 0; i < len(m.books); i++ {
		if m.books[i].ID == id {
			m.books = append(m.books[:i], m.books[i+1:]...)
			return "this book ID is deleted", nil
		}
	}
	return "", errors.New("book not found")
}
