package main

import (
	"context"
	"fmt"
	"gin/handler"
	gdb "gin/repository/gorm"

	// repository "gin/repository/mysql"
	rdb "gin/repository/redis"
	"gin/service"
	"time"

	// "log"
	"database/sql"
	"encoding/json"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var CTX = context.Background()

type BookRedisTest struct {
	Key        string
	Value      string
	Expiration time.Duration
}

func main() {
	initConfig()
	// r := setupRoute()
	// r.Run(viper.GetString("app.port"))
	db := initGorm()
	bookRepoGORM := gdb.NewRepositoryGORM(db)
	result, err := bookRepoGORM.GetBookAuthor(1)
	if err != nil {
		fmt.Println(err)
	}
	js, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", js)
	result2, err := bookRepoGORM.GetAuthorBook(1)
	if err != nil {
		fmt.Println(result2)
	}

	js, err = json.Marshal(result2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("json %s\n", js)

}

func setupRoute() *gin.Engine {
	r := gin.Default()

	db := initGorm()
	rd := initRedis()
	// _ = db
	// _ = rd
	bookRepoRedis := rdb.NewBookRepositoryRedis(rd)
	// bookRepo := repository.NewRepositoryDB(db)
	bookRepo := gdb.NewRepositoryGORM(db)
	bookService := service.NewBookService(bookRepo, bookRepoRedis)
	bookHandler := handler.NewBookHandler(bookService)

	r.GET("/books", bookHandler.GetAllBook)
	r.GET("/books/:id", bookHandler.GetByID)
	r.GET("/books/:id/name", bookHandler.GetBookAuthor)
	r.GET("/books/name/:id", bookHandler.GetAuthorBook)
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

func initGorm() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.hostname"),
		viper.GetInt("db.port"),
		viper.GetString("db.dbname"),
	)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&gdb.Books{})
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
