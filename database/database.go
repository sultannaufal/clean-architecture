package database

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sultannaufal/clean-architecture/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Rdb *redis.Client
var books = []model.Book{
	{ID: 1, Title: "Go Tutorial", Isbn: strconv.Itoa(model.Isbn), Writer: "Anon", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

// var rDB

func CreateConnection() {

	var e error
	DB, e = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func CreateRedisConnection() {
	config := map[string]string{
		"REDIS_Username": os.Getenv("REDIS_USERNAME"),
		"REDIS_Password": os.Getenv("REDIS_PASSWORD"),
		"REDIS_Port":     os.Getenv("REDIS_PORT"),
		"REDIS_Host":     os.Getenv("REDIS_HOST"),
		"REDIS_Name":     os.Getenv("REDIS_NAME"),
	}

	connectionString := fmt.Sprintf("redis://%s:%s@%s:%s/%s",
		config["REDIS_Username"],
		config["REDIS_Password"],
		config["REDIS_Host"],
		config["REDIS_Port"],
		config["REDIS_Name"])

	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		panic(err)
	}

	Rdb = redis.NewClient(opt)
	ttl := time.Duration(3) * time.Hour

	book_json, _ := json.Marshal(books)

	// store data
	op1 := Rdb.Set(context.Background(), "book", string(book_json), ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{})
}

func GetConnection() *gorm.DB {
	if DB == nil {
		CreateConnection()
	}
	if Rdb == nil {
		CreateRedisConnection()
	}
	return DB
}
