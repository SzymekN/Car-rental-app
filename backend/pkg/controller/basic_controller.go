package controller

import (
	"github.com/labstack/echo/v4"
)

type BasicHandler interface {
	RegisterRoutes()
	Save(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetById(c echo.Context) error
	GetAll(c echo.Context) error
}
