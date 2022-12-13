package controller

import (
	"fmt"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/executor"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"

	"github.com/labstack/echo/v4"
)

type RepairHandler struct {
	sysOperator producer.SystemOperator
	authConf    auth.AuthConfig
	group       *echo.Group
}

func NewRepairHandler(sysOp producer.SystemOperator, ac auth.AuthConfig, g *echo.Group) *RepairHandler {
	uh := &RepairHandler{
		sysOperator: sysOp,
		group:       g,
		authConf:    ac,
	}
	fmt.Println(sysOp)
	return uh
}

func (uh *RepairHandler) RegisterRoutes() {
	uh.group.GET("/repairs", uh.GetById, uh.authConf.IsAuthorized)
	uh.group.GET("/repairs/all", uh.GetAll, uh.authConf.IsAuthorized)
	uh.group.POST("/repairs", uh.Save, uh.authConf.IsAuthorized)
	uh.group.PUT("/repairs", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/repairs", uh.Delete, uh.authConf.IsAuthorized)
}

func (uh *RepairHandler) Save(c echo.Context) error {
	d, l := executor.GenericPost(c, uh.sysOperator, model.Repair{})
	return HandleRequestResult(c, d, l)
}

func (uh *RepairHandler) Update(c echo.Context) error {
	d, l := executor.GenericUpdate(c, uh.sysOperator, model.Repair{})
	return HandleRequestResult(c, d, l)
}

func (uh *RepairHandler) Delete(c echo.Context) error {
	d, l := executor.GenericDelete(c, uh.sysOperator, model.Repair{})
	return HandleRequestResult(c, d, l)
}

func (uh *RepairHandler) GetById(c echo.Context) error {
	d, l := executor.GenericGetById(c, uh.sysOperator, model.Repair{})
	return HandleRequestResult(c, d, l)
}

func (uh *RepairHandler) GetAll(c echo.Context) error {
	d, l := executor.GenericGetAll(c, uh.sysOperator, []model.Repair{})
	return HandleRequestResult(c, d, l)
}
