package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

	// var validToken string
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("SaveEmployee ")

	// user to save in the database
	var me model.Employee

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	// try saving data got in the request to the User datatype
	me, logger.Log = executor.BindData(c, me)
	if logger.Log.Err != nil {
		return logger.Log.Err
	}

	// hash password
	me.User.Password, logger.Log = auth.GeneratehashPassword(me.User.Password)
	if logger.Err != nil {
		return logger.Err
	}

	db := uh.sysOperator.DB
	result := db.Model(&model.Employee{}).Preload("User").Debug().Create(&me)
	if err := result.Error; err != nil {
		code := http.StatusBadRequest
		msg := fmt.Sprintf("[ERROR]: couldn't create user, HTTP: %v", code)
		logger.Log.Populate("err", msg, code, err)
	}

	logger.Log = executor.CheckIfAffected(result)

	if logger.Err != nil {
		if logger.Err.Error() == "no rows affected" {
			code := http.StatusBadRequest
			msg := fmt.Sprintf("[ERROR]: duplicate entry - user exists, HTTP: %v", code)
			err := errors.New("duplicate entry")
			logger.Log.Populate("err", msg, code, err)
			return logger.Err
		} else {
			return logger.Err
		}
	}

	code := http.StatusOK
	k := "info"
	msg := "[INFO] Employee insert completed: user signed up, email: {" + me.User.Email + "}, HTTP: " + strconv.Itoa(code)
	logger.Populate(k, msg, code, nil)
	return c.JSON(code, msg)

}

func (uh *EmployeeHandler) Update(c echo.Context) error {
	d, l := executor.GenericUpdate(c, uh.sysOperator, model.Employee{})
	return HandleRequestResult(c, d, l)
}

func (uh *EmployeeHandler) Delete(c echo.Context) error {
	d, l := executor.GenericDelete(c, uh.sysOperator, model.Employee{})
	return HandleRequestResult(c, d, l)
}

func (uh *EmployeeHandler) GetById(c echo.Context) error {
	d, l := executor.GenericGetById(c, uh.sysOperator, model.Employee{})
	return HandleRequestResult(c, d, l)
}

func (uh *EmployeeHandler) GetAll(c echo.Context) error {
	d, l := executor.GenericGetAll(c, uh.sysOperator, []model.Employee{})
	return HandleRequestResult(c, d, l)
}
