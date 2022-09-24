package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sultannaufal/clean-architecture/database"
	"github.com/sultannaufal/clean-architecture/internal/http"
	"github.com/sultannaufal/clean-architecture/internal/middleware"
)

func main() {
	database.CreateConnection()
	database.CreateRedisConnection()
	e := echo.New()
	middleware.Init(e)
	http.NewHttp(e)
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
