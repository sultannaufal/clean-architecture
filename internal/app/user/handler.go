package user

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sultannaufal/clean-architecture/database"
	"github.com/sultannaufal/clean-architecture/internal/dto"
	"github.com/sultannaufal/clean-architecture/internal/model"
	"github.com/sultannaufal/clean-architecture/internal/repository"
	"github.com/sultannaufal/clean-architecture/pkg/util/hash"
	. "github.com/sultannaufal/clean-architecture/pkg/util/response"
)

func Get(c echo.Context) error {

	payload := new(dto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	result, err := repository.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, Response{Code: http.StatusOK, Message: "success", Data: result})
}

func GetByID(c echo.Context) error {

	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	if err := c.Validate(payload); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	result, err := repository.FindByID(payload.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, Response{Code: http.StatusOK, Message: "success", Data: result})
}

func Create(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	hashedPassword, err := hash.Make(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	user.Password = hashedPassword

	if err := database.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, Response{Code: http.StatusOK, Message: "Success", Data: user})
}

func Update(c echo.Context) error {
	auth := c.Get("authInfo").(jwt.MapClaims)
	auth_id := fmt.Sprintf("%v", auth["id"])
	id := c.Param("id")
	if auth_id != id {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusInternalServerError, Message: "Cannot delete other user"})
	}
	result := database.DB.Model(&model.User{}).Where("id = ?", id)
	if err := result.Error; err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	hashedPassword, err := hash.Make(c.FormValue("password"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	update := result.Updates(model.User{Name: c.FormValue("name"), Email: c.FormValue("email"), Password: hashedPassword})
	if err := update.Error; err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, Response{Code: http.StatusOK, Message: "Success"})
}

func Delete(c echo.Context) error {
	auth := c.Get("authInfo").(jwt.MapClaims)
	auth_id := fmt.Sprintf("%v", auth["id"])
	id := c.Param("id")
	if auth_id != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot update other user")
	}
	result := database.DB.Where("id = ?", id).Delete(&model.User{})
	if err := result.Error; err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, Response{Code: http.StatusOK, Message: "Success"})
}
