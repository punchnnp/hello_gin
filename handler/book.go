package handler

import (
	"fmt"
	"gin/service"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

type UserClaim struct {
	jwt.RegisteredClaims
	Name      string
	ExpiresAt int
}

func LoginHandler(c *gin.Context) {
	signature := []byte("flowers")
	id := c.Param("name")
	// expire should be time format unix for any timezone
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		Name: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(5 * time.Minute)},
		},
	})

	ss, err := token.SignedString(signature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}

func validateToken(token string) error {
	newToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("flowers"), nil
	})

	// check if token valid
	_, ok := newToken.Claims.(jwt.Claims)
	if !ok || !newToken.Valid {
		return err
	}

	// check if expire or not numeric time
	time, err := newToken.Claims.GetExpirationTime()
	_ = time

	// check if expire
	return err
}

func AuthoruzationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	ss := strings.TrimPrefix("Bearer ", s)

	err := validateToken(ss)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
