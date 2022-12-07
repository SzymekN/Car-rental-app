package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// nie wiem co to ma robić ale to chyba było do przechowywania wszystkich HandlerObjects i generycznej rejestracji
type SystemController struct {
	Controllers []BasicController
}

type BasicController interface {
	RegisterRoutes(e echo.Group, db *gorm.DB)
	Save(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetById(c echo.Context) error
	GetAll(c echo.Context) error
}
