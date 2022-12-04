package controller

import (
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	db *gorm.DB
}

type Registrator interface {
	RegisterRoutes(e echo.Group, db *gorm.DB)
}

type UsersHandler interface {
	SaveUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	GetUserById(c echo.Context) error
	GetUsers(c echo.Context) error
}

func (uc *UsersController) SaveUser(c echo.Context) error {
	return GenericPost(c, uc.db, model.User{})
}

func (uc *UsersController) UpdateUser(c echo.Context) error {
	return GenericUpdate(c, uc.db, model.User{})
}

func (uc *UsersController) DeleteUser(c echo.Context) error {
	return GenericDelete(c, uc.db, model.User{})
}

func (uc *UsersController) GetUserById(c echo.Context) error {
	return GenericGetById(c, uc.db, model.User{})
}

func (uc *UsersController) GetUsers(c echo.Context) error {
	return GenericGetAll(c, uc.db, []model.User{})
}
