package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sultannaufal/clean-architecture/pkg/util/log"
)

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.MakeLogEntry(c).Info("incoming request")
		return next(c)
	}
}
