package controller

import (
	"fmt"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/executor"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"

	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	sysOperator producer.SystemOperator
	authConf    auth.AuthConfig
	group       *echo.Group
}

func NewNotificationHandler(sysOp producer.SystemOperator, ac auth.AuthConfig, g *echo.Group) *NotificationHandler {
	uh := &NotificationHandler{
		sysOperator: sysOp,
		group:       g,
		authConf:    ac,
	}
	fmt.Println(sysOp)
	return uh
}

func (uh *NotificationHandler) RegisterRoutes() {
	uh.group.GET("/notifications", uh.GetById, uh.authConf.IsAuthorized)
	uh.group.GET("/notifications/all", uh.GetAll, uh.authConf.IsAuthorized)
	uh.group.POST("/notifications", uh.Save, uh.authConf.IsAuthorized)
	uh.group.PUT("/notifications", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/notifications", uh.Delete, uh.authConf.IsAuthorized)
}

func (uh *NotificationHandler) Save(c echo.Context) error {
	return executor.GenericPost(c, uh.sysOperator, model.Notification{})
}

func (uh *NotificationHandler) Update(c echo.Context) error {
	return executor.GenericUpdate(c, uh.sysOperator, model.Notification{})
}

func (uh *NotificationHandler) Delete(c echo.Context) error {
	return executor.GenericDelete(c, uh.sysOperator, model.Notification{})
}

func (uh *NotificationHandler) GetById(c echo.Context) error {
	return executor.GenericGetById(c, uh.sysOperator, model.Notification{})
}

func (uh *NotificationHandler) GetAll(c echo.Context) error {
	return executor.GenericGetAll(c, uh.sysOperator, []model.Notification{})
}
