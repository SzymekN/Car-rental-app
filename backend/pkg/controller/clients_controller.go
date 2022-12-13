package controller

import (
	"fmt"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/executor"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"

	"github.com/labstack/echo/v4"
)

type ClientHandler struct {
	sysOperator producer.SystemOperator
	authConf    auth.AuthConfig
	group       *echo.Group
}

func NewClientHandler(sysOp producer.SystemOperator, ac auth.AuthConfig, g *echo.Group) *ClientHandler {
	uh := &ClientHandler{
		sysOperator: sysOp,
		group:       g,
		authConf:    ac,
	}
	fmt.Println(sysOp)
	return uh
}

func (uh *ClientHandler) RegisterRoutes() {
	uh.group.GET("/clients", uh.GetById, uh.authConf.IsAuthorized)
	uh.group.GET("/clients/all", uh.GetAll, uh.authConf.IsAuthorized)
	uh.group.POST("/clients", uh.Save, uh.authConf.IsAuthorized)
	uh.group.PUT("/clients", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/clients", uh.Delete, uh.authConf.IsAuthorized)
	uh.group.PUT("/clients/self", uh.UpdateSelf, uh.authConf.IsAuthorized)
	uh.group.DELETE("/clients/self", uh.DeleteSelf, uh.authConf.IsAuthorized)
}

func (uh *ClientHandler) Save(c echo.Context) error {
	return executor.GenericPost(c, uh.sysOperator, model.Client{})
}

func (uh *ClientHandler) Update(c echo.Context) error {
	return executor.GenericUpdate(c, uh.sysOperator, model.Client{})
}

func (uh *ClientHandler) Delete(c echo.Context) error {
	return executor.GenericDelete(c, uh.sysOperator, model.Client{})
}

func (uh *ClientHandler) GetById(c echo.Context) error {
	return executor.GenericGetById(c, uh.sysOperator, model.Client{})
}

func (uh *ClientHandler) GetAll(c echo.Context) error {
	return executor.GenericGetAll(c, uh.sysOperator, []model.Client{})
}
func (uh *ClientHandler) UpdateSelf(c echo.Context) error {
	return executor.GenericUpdate(c, uh.sysOperator, model.Client{})
}

func (uh *ClientHandler) DeleteSelf(c echo.Context) error {
	return executor.GenericDelete(c, uh.sysOperator, model.Client{})
}
