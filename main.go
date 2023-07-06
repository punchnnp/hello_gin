package main

import (
	"context"
	"fmt"
	"gin/handler"
	repository "gin/repository/mysql"
	rdb "gin/repository/redis"
	"gin/service"
	"time"

	// "log"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var CTX = context.Background()

type BookRedisTest struct {
	Key        string
	Value      string
	Expiration time.Duration
}

func main() {
	initConfig()
	r := setupRoute()

	// db := initDB()
	// bookRepo := repository.NewRepositoryDB(db)

	rd := initRedis()
	bookRepoRedis := rdb.NewBookRepositoryRedis(rd)

	test := rdb.BookRedis{
		Key:        1,
		Value:      "Nunnapat",
		Expiration: 0,
	}
	err := bookRepoRedis.Set(test)
	if err != nil {
		panic(err)
	}
	val, err2 := bookRepoRedis.Get(1)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(val)
	r.Run(viper.GetString("app.port"))
}

func setupRoute() *gin.Engine {
	r := gin.Default()

	db := initDB()
	rd := initRedis()
	_ = db
	_ = rd
	bookRepoRedis := rdb.NewBookRepositoryRedis(rd)
	bookRepo := repository.NewBookRepositoryMock()
	bookService := service.NewBookService(bookRepo, bookRepoRedis)
	bookHandler := handler.NewBookHandler(bookService)

	r.GET("/books", bookHandler.GetAllBook)
	r.GET("/books/:id", bookHandler.GetByID)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.POST("/books", bookHandler.AddBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

	return r
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDB() *sql.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.hostname"),
		viper.GetInt("db.port"),
		viper.GetString("db.dbname"),
	)

	db, err := sql.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db
}

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	return rdb
}
