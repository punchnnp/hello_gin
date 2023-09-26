package main

import (
	"context"
	"fmt"
	"gin/handler"
	"gin/middleware"
	gdb "gin/repository/gorm"

	rdb "gin/repository/redis"
	"gin/service"
	"time"

	// "log"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var CTX = context.Background()

func main() {
	initConfig()
	r := setupRoute()
	r.Run(viper.GetString("app.port"))
	// db := initGorm()
	// bookRepo := gdb.NewRepositoryGORM(db)
	// auts, err := bookRepo.GetAutByID(2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(auts)
}

func setupRoute() *gin.Engine {
	r := gin.Default()

	r.POST("/login/:name", handler.LoginHandler)
	db := initGorm()
	dbb := initDB()
	rd := initRedis()
	_ = dbb
	// _ = db
	// _ = rd
	bookRepoRedis := rdb.NewBookRepositoryRedis(rd)
	// bookRepo := repository.NewRepositoryDB(db)
	bookRepo := gdb.NewRepositoryGORM(db)
	bookService := service.NewBookService(bookRepo, bookRepoRedis)
	bookHandler := handler.NewBookHandler(bookService)

	groups := r.Group("/", middleware.AuthoruzationMiddleware)

	groups.GET("/books", bookHandler.GetAllBook)
	groups.GET("/books/:id", bookHandler.GetByID)
	groups.GET("/books/:id/name", bookHandler.GetBookAuthor)
	groups.GET("/books/name/:id", bookHandler.GetAuthorBook)
	groups.PUT("/books/:id", bookHandler.UpdateBook)
	groups.POST("/books", bookHandler.AddBook)
	groups.DELETE("/books/:id", bookHandler.DeleteBook)

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
