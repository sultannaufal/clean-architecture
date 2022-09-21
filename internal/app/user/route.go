package user

import (
	"github.com/labstack/echo/v4"
	"github.com/sultannaufal/clean-architecture/internal/middleware"
)

func Route(g *echo.Group) {
	g.GET("", Get, middleware.AuthMiddleware)
	g.GET("/:id", GetByID, middleware.AuthMiddleware)
	g.POST("", Create)
	g.PUT("/:id", Update, middleware.AuthMiddleware)
	g.DELETE("/:id", Delete, middleware.AuthMiddleware)
}
