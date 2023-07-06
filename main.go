package main

import (
	"context"
	"fmt"
	"gin/handler"
	"gin/repository"
	"gin/service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var CTX = context.Background()

func hello(c *gin.Context) {
	c.String(200, "hello world")
}

func multiply(c *gin.Context) {
	input, err := strconv.ParseInt(c.Param("num"), 10, 32)
	if err != nil {
		log.Fatalf(err.Error())
	}
	result := input * 3
	c.String(200, "%d multiply by 3 equals to %d", input, result)
}

func setupRoute(r *gin.Engine) {
	// r := gin.Default()

	bookRepo := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to my GIN practice")
	})
	r.GET("/math/:num", multiply)
	r.GET("/hello", hello)
	r.GET("/books", bookHandler.GetAllBook)
	r.GET("/books/:id", bookHandler.GetByID)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.POST("/books", bookHandler.AddBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

	// return r
}

func DummyMiddleWare() gin.HandlerFunc {
	fmt.Println("I'm Dummy")

	return func(c *gin.Context) {
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(DummyMiddleWare())
	setupRoute(r)
	fmt.Println("testing")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(CTX, "name", "Punch", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err2 := rdb.Get(CTX, "name").Result()
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("Redis start %v", val)

	r.Run(":3000")
}
