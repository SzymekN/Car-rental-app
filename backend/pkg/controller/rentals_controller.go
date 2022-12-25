package controller

import (
	"fmt"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/executor"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"

	"github.com/labstack/echo/v4"
)

type RentalHandler struct {
	sysOperator producer.SystemOperator
	authConf    auth.AuthConfig
	group       *echo.Group
}

func NewRentalHandler(sysOp producer.SystemOperator, ac auth.AuthConfig, g *echo.Group) *RentalHandler {
	uh := &RentalHandler{
		sysOperator: sysOp,
		group:       g,
		authConf:    ac,
	}
	fmt.Println(sysOp)
	return uh
}

func (uh *RentalHandler) RegisterRoutes() {
	uh.group.GET("/rentals", uh.GetById, uh.authConf.IsAuthorized)
	uh.group.GET("/rentals/all", uh.GetAll, uh.authConf.IsAuthorized)
	uh.group.POST("/rentals", uh.Save, uh.authConf.IsAuthorized)
	uh.group.PUT("/rentals", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/rentals", uh.Delete, uh.authConf.IsAuthorized)
	uh.group.POST("/rentals/self", uh.SaveSelf, uh.authConf.IsAuthorized)
}

func (uh *RentalHandler) Save(c echo.Context) error {
	d, l := executor.GenericPost(c, uh.sysOperator, model.Rental{})
	return HandleRequestResult(c, d, l)
}

func (uh *RentalHandler) SaveSelf(c echo.Context) error {
	mr := model.Rental{}
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("SelfRental ")

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	mr, logger.Log = executor.BindData(c, mr)
	if logger.Err != nil {
		return logger.Err
	}

	var uid int
	uid, logger.Log = GetUIDFromContextToken(c, uh.sysOperator)
	if logger.Err != nil {
		return logger.Err
	}
	mr.ClientID = uid

	d, l := executor.GenericPost(c, uh.sysOperator, mr)
	if logger.Err != nil {
		return logger.Err
	}
	return c.JSON(l.Code, d)
}

func (uh *RentalHandler) Update(c echo.Context) error {
	d, l := executor.GenericUpdate(c, uh.sysOperator, model.Rental{})
	return HandleRequestResult(c, d, l)
}

func (uh *RentalHandler) Delete(c echo.Context) error {
	d, l := executor.GenericDelete(c, uh.sysOperator, model.Rental{})
	return HandleRequestResult(c, d, l)
}

func (uh *RentalHandler) GetById(c echo.Context) error {
	d, l := executor.GenericGetById(c, uh.sysOperator, model.Rental{})
	return HandleRequestResult(c, d, l)
}

func (uh *RentalHandler) GetAll(c echo.Context) error {
	d, l := executor.GenericGetAll(c, uh.sysOperator, []model.Rental{})
	return HandleRequestResult(c, d, l)
}
