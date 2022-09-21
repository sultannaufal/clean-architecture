package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sultannaufal/clean-architecture/pkg/util/validator"
)

func Init(e *echo.Echo) {
	e.Use(
		echoMiddleware.Recover(),
		echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),
	)
	e.Use(LogMiddleware)
	e.HTTPErrorHandler = ErrorHandler
	e.Validator = validator.Validator

}
