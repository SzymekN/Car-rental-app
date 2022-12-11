package controller

import (
	"fmt"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/executor"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"

	"github.com/labstack/echo/v4"
)

type VehicleHandler struct {
	sysOperator producer.SystemOperator
	authConf    auth.AuthConfig
	group       *echo.Group
}

func NewVehicleHandler(sysOp producer.SystemOperator, ac auth.AuthConfig, g *echo.Group) *VehicleHandler {
	uh := &VehicleHandler{
		sysOperator: sysOp,
		group:       g,
		authConf:    ac,
	}
	fmt.Println(sysOp)
	return uh
}

func (uh *VehicleHandler) RegisterRoutes() {
	uh.group.GET("/vehicles", uh.GetById)
	uh.group.GET("/vehicles/all", uh.GetAll)
	uh.group.POST("/vehicles", uh.Save, uh.authConf.IsAuthorized)
	uh.group.PUT("/vehicles", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/vehicles", uh.Delete, uh.authConf.IsAuthorized)
}

func (uh *VehicleHandler) Save(c echo.Context) error {
	return executor.GenericPost(c, uh.sysOperator, model.Vehicle{})
}

func (uh *VehicleHandler) Update(c echo.Context) error {
	return executor.GenericUpdate(c, uh.sysOperator, model.Vehicle{})
}

func (uh *VehicleHandler) Delete(c echo.Context) error {
	return executor.GenericDelete(c, uh.sysOperator, model.Vehicle{})
}

func (uh *VehicleHandler) GetById(c echo.Context) error {
	return executor.GenericGetById(c, uh.sysOperator, model.Vehicle{})
}

func (uh *VehicleHandler) GetAll(c echo.Context) error {
	return executor.GenericGetAll(c, uh.sysOperator, []model.Vehicle{})
}
