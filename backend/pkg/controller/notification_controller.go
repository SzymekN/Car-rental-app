package controller

import (
	"fmt"
	"net/http"

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
	uh.group.POST("/notifications/info", uh.GetById, uh.authConf.IsAuthorized)
	uh.group.GET("/notifications/all", uh.GetAll, uh.authConf.IsAuthorized)
	uh.group.POST("/notifications", uh.Save, uh.authConf.IsAuthorized)
	uh.group.POST("/notifications/employee", uh.AddByEmployee, uh.authConf.IsAuthorized)
	uh.group.POST("/notifications/client", uh.AddByClient, uh.authConf.IsAuthorized)
	uh.group.PUT("/notifications", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/notifications", uh.Delete, uh.authConf.IsAuthorized)
}

func (uh *NotificationHandler) Save(c echo.Context) error {
	d, l := executor.GenericPost(c, uh.sysOperator, model.Notification{})
	return HandleRequestResult(c, d, l)
}

func (uh *NotificationHandler) AddByEmployee(c echo.Context) error {
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("AddNotificationByEmployee ")
	db := uh.sysOperator.GetDB()
	n := model.Notification{}

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	var eid int
	eid, logger.Log = GetEIDFromContextToken(c, uh.sysOperator)
	if logger.Err != nil {
		return logger.Err
	}

	n, logger.Log = executor.BindData(c, n)
	if logger.Err != nil && n.GetId() < 0 {
		return logger.Err
	}

	n.EmployeeID = &eid
	logger.Log = executor.Insert(c, db, n)
	if logger.Log.Err != nil {
		return logger.Err
	}
	logger.Log.Code = http.StatusOK
	logger.Log.Key = "info"
	logger.Log.Msg = fmt.Sprintf("[INFO] completed, HTTP: %v", logger.Log.Code)
	return c.JSON(logger.Code, n)
	// d, l := executor.GenericPost(c, uh.sysOperator, model.Notification{})
	// return HandleRequestResult(c, d, l)
}

func (uh *NotificationHandler) AddByClient(c echo.Context) error {
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("AddNotificationByClient ")
	db := uh.sysOperator.GetDB()
	n := model.Notification{}

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	var cid int
	cid, logger.Log = GetCIDFromContextToken(c, uh.sysOperator)
	if logger.Err != nil {
		return logger.Err
	}

	n, logger.Log = executor.BindData(c, n)
	if logger.Err != nil && n.GetId() < 0 {
		return logger.Err
	}

	n.ClientID = &cid
	logger.Log = executor.Insert(c, db, n)
	if logger.Log.Err != nil {
		return logger.Err
	}
	logger.Log.Code = http.StatusOK
	logger.Log.Key = "info"
	logger.Log.Msg = fmt.Sprintf("[INFO] completed, HTTP: %v", logger.Log.Code)
	return c.JSON(logger.Code, n)
}

func (uh *NotificationHandler) Update(c echo.Context) error {
	d, l := executor.GenericUpdate(c, uh.sysOperator, model.Notification{})
	return HandleRequestResult(c, d, l)
}

func (uh *NotificationHandler) Delete(c echo.Context) error {
	d, l := executor.GenericDelete(c, uh.sysOperator, model.Notification{})
	return HandleRequestResult(c, d, l)
}

func (uh *NotificationHandler) GetById(c echo.Context) error {
	d, l := executor.GenericGetById(c, uh.sysOperator, model.Notification{})
	return HandleRequestResult(c, d, l)
}

func (uh *NotificationHandler) GetAll(c echo.Context) error {
	d, l := executor.GenericGetAll(c, uh.sysOperator, []model.Notification{})
	return HandleRequestResult(c, d, l)
}
