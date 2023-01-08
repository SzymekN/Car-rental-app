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
	uh.group.POST("/salaries/info", uh.GetById, uh.authConf.IsAuthorized)
	uh.group.GET("/salaries/all", uh.GetAll, uh.authConf.IsAuthorized)
	uh.group.POST("/salaries", uh.Save, uh.authConf.IsAuthorized)
	uh.group.PUT("/salaries", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/salaries", uh.Delete, uh.authConf.IsAuthorized)
}

func (uh *SalaryHandler) Save(c echo.Context) error {
	d, l := executor.GenericPost(c, uh.sysOperator, model.Salary{})
	return HandleRequestResult(c, d, l)
}

func (uh *SalaryHandler) Update(c echo.Context) error {
	d, l := executor.GenericUpdate(c, uh.sysOperator, model.Salary{})
	return HandleRequestResult(c, d, l)
}

func (uh *SalaryHandler) Delete(c echo.Context) error {
	d, l := executor.GenericDelete(c, uh.sysOperator, model.Salary{})
	return HandleRequestResult(c, d, l)
}

func (uh *SalaryHandler) GetById(c echo.Context) error {
	d, l := executor.GenericGetById(c, uh.sysOperator, model.Salary{})
	return HandleRequestResult(c, d, l)
}

func (uh *SalaryHandler) GetAll(c echo.Context) error {
	d, l := executor.GenericGetAll(c, uh.sysOperator, []model.Salary{})
	return HandleRequestResult(c, d, l)
}
