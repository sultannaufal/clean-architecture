package http

import (
	"github.com/labstack/echo/v4"
	"github.com/sultannaufal/clean-architecture/internal/app/auth"
	"github.com/sultannaufal/clean-architecture/internal/app/book"
	"github.com/sultannaufal/clean-architecture/internal/app/user"
)

func NewHttp(e *echo.Echo) {
	api := e.Group("/api")
	v1 := api.Group("/v1")
	user.Route(v1.Group("/users"))
	auth.Route(v1.Group("/auth"))
	book.Route(v1.Group("/books"))
}
