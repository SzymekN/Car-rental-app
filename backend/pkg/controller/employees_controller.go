package controller

import (
	"fmt"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/executor"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"

	"github.com/labstack/echo/v4"
)

type EmployeeHandler struct {
	sysOperator producer.SystemOperator
	authConf    auth.AuthConfig
	group       *echo.Group
}

func NewEmployeeHandler(sysOp producer.SystemOperator, ac auth.AuthConfig, g *echo.Group) *EmployeeHandler {
	uh := &EmployeeHandler{
		sysOperator: sysOp,
		group:       g,
		authConf:    ac,
	}
	fmt.Println(sysOp)
	return uh
}

func (uh *EmployeeHandler) RegisterRoutes() {
	uh.group.GET("/employees", uh.GetById, uh.authConf.IsAuthorized)
	uh.group.GET("/employees/all", uh.GetAll, uh.authConf.IsAuthorized)
	uh.group.POST("/employees", uh.Save, uh.authConf.IsAuthorized)
	uh.group.PUT("/employees", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/employees", uh.Delete, uh.authConf.IsAuthorized)
}

func (uh *EmployeeHandler) Save(c echo.Context) error {
	return executor.GenericPost(c, uh.sysOperator, model.Employee{})
}

func (uh *EmployeeHandler) Update(c echo.Context) error {
	return executor.GenericUpdate(c, uh.sysOperator, model.Employee{})
}

func (uh *EmployeeHandler) Delete(c echo.Context) error {
	return executor.GenericDelete(c, uh.sysOperator, model.Employee{})
}

func (uh *EmployeeHandler) GetById(c echo.Context) error {
	return executor.GenericGetById(c, uh.sysOperator, model.Employee{})
}

func (uh *EmployeeHandler) GetAll(c echo.Context) error {
	return executor.GenericGetAll(c, uh.sysOperator, []model.Employee{})
}
