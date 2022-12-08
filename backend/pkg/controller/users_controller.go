package controller

import (
	"fmt"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/executor"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	sysOperator producer.SystemOperator
	authConf    auth.AuthConfig
	group       *echo.Group
}

func NewUserHandler(sysOp producer.SystemOperator, ac auth.AuthConfig, g *echo.Group) UserHandler {
	uh := UserHandler{
		sysOperator: sysOp,
		group:       g,
		authConf:    ac,
	}
	fmt.Println(sysOp)
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
	return executor.GenericPost(c, uh.sysOperator, model.User{})
}

func (uh *UserHandler) Update(c echo.Context) error {
	return executor.GenericUpdate(c, uh.sysOperator, model.User{})
}

func (uh *UserHandler) Delete(c echo.Context) error {
	return executor.GenericDelete(c, uh.sysOperator, model.User{})
}

func (uh *UserHandler) GetById(c echo.Context) error {
	return executor.GenericGetById(c, uh.sysOperator, model.User{})
}

func (uh *UserHandler) GetAll(c echo.Context) error {
	cos := executor.GenericGetAll(c, uh.sysOperator, []model.User{})
	return cos
}
