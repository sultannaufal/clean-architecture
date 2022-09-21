package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sultannaufal/clean-architecture/database"
	"github.com/sultannaufal/clean-architecture/internal/http"
	"github.com/sultannaufal/clean-architecture/internal/middleware"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	database.CreateConnection()
	database.CreateRedisConnection()
	e := echo.New()
	middleware.Init(e)
	http.NewHttp(e)
	e.Logger.Fatal(e.Start(":8000"))
}
