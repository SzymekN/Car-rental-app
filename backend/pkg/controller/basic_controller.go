package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BasicController interface {
	RegisterRoutes(e echo.Group, db *gorm.DB)
	Save(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetById(c echo.Context) error
	GetAll(c echo.Context) error
}
