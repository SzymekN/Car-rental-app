package controller

import (
	"fmt"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/executor"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"

	"github.com/labstack/echo/v4"
)

type SalaryHandler struct {
	sysOperator producer.SystemOperator
	authConf    auth.AuthConfig
	group       *echo.Group
}

func NewSalaryHandler(sysOp producer.SystemOperator, ac auth.AuthConfig, g *echo.Group) *SalaryHandler {
	uh := &SalaryHandler{
		sysOperator: sysOp,
		group:       g,
		authConf:    ac,
	}
	fmt.Println(sysOp)
	return uh
}

func (uh *SalaryHandler) RegisterRoutes() {
	uh.group.GET("/salaries", uh.GetById, uh.authConf.IsAuthorized)
	uh.group.GET("/salaries/all", uh.GetAll, uh.authConf.IsAuthorized)
	uh.group.POST("/salaries", uh.Save, uh.authConf.IsAuthorized)
	uh.group.PUT("/salaries", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/salaries", uh.Delete, uh.authConf.IsAuthorized)
}

func (uh *SalaryHandler) Save(c echo.Context) error {
	return executor.GenericPost(c, uh.sysOperator, model.Salary{})
}

func (uh *SalaryHandler) Update(c echo.Context) error {
	return executor.GenericUpdate(c, uh.sysOperator, model.Salary{})
}

func (uh *SalaryHandler) Delete(c echo.Context) error {
	return executor.GenericDelete(c, uh.sysOperator, model.Salary{})
}

func (uh *SalaryHandler) GetById(c echo.Context) error {
	return executor.GenericGetById(c, uh.sysOperator, model.Salary{})
}

func (uh *SalaryHandler) GetAll(c echo.Context) error {
	return executor.GenericGetAll(c, uh.sysOperator, []model.Salary{})
}
