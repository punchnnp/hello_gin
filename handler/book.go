package handler

import (
	"gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) bookHandler {
	return bookHandler{bookService: bookService}
}

func (h bookHandler) GetAllBook(c *gin.Context) {
	books, err := h.bookService.GetAllBook()
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h bookHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(err)
		return
	}
	book, err2 := h.bookService.GetByID(int(id))
	if err2 != nil {
		// http status depend on specification
		c.String(http.StatusOK, err2.Error())
		return
	}
	c.JSON(http.StatusOK, book)

}

func (h bookHandler) GetBookAuthor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(err)
		return
	}
	aut, err2 := h.bookService.GetBookAuthor(int(id))
	if err2 != nil {
		c.String(http.StatusOK, err2.Error())
		return
	}
	c.JSON(http.StatusOK, aut)
}

func (h bookHandler) GetAuthorBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(err)
		return
	}
	books, err2 := h.bookService.GetAuthorBook(int(id))
	if err != nil {
		c.String(http.StatusOK, err2.Error())
	}
	c.JSON(http.StatusOK, books)
}

func (h bookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(err)
		return
	}
	book, err2 := h.bookService.UpdateBook(int(id))
	if err2 != nil {
		c.String(http.StatusOK, err2.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h bookHandler) AddBook(c *gin.Context) {
	book, err := h.bookService.AddBook()
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h bookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(err)
		return
	}

	result, err2 := h.bookService.DeleteBook(int(id))
	if err2 != nil {
		c.String(http.StatusOK, err2.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}
