package controller

import (
	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	SystemOperator
	authConf auth.AuthConfig
	group    *echo.Group
}

func NewUserHandler(db *gorm.DB, l model.SystemLogger, ac auth.AuthConfig, g *echo.Group) UserHandler {
	uh := UserHandler{
		SystemOperator: SystemOperator{
			DB:           db,
			SystemLogger: l,
		},
		authConf: ac,
		group:    g,
	}
	return uh
}

func (uh *UserHandler) RegisterRoutes() {
	uh.group.GET("/users", uh.GetById)
	uh.group.GET("/users/all", uh.GetAll)
	uh.group.POST("/users", uh.Save, uh.authConf.IsAuthorized)
	uh.group.PUT("/users", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/users", uh.Delete, uh.authConf.IsAuthorized)
}

func (uh *UserHandler) Save(c echo.Context) error {
	return GenericPost(c, uh.SystemOperator, model.User{})
}

func (uh *UserHandler) Update(c echo.Context) error {
	return GenericUpdate(c, uh.SystemOperator, model.User{})
}

func (uh *UserHandler) Delete(c echo.Context) error {
	return GenericDelete(c, uh.SystemOperator, model.User{})
}

func (uh *UserHandler) GetById(c echo.Context) error {
	return GenericGetById(c, uh.SystemOperator, model.User{})
}

func (uh *UserHandler) GetAll(c echo.Context) error {
	return GenericGetAll(c, uh.SystemOperator, []model.User{})
}
