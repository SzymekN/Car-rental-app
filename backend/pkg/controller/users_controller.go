package controller

import (
	"fmt"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/executor"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	sysOperator producer.SystemOperator
	authConf    auth.AuthConfig
	group       *echo.Group
}

func NewUserHandler(sysOp producer.SystemOperator, ac auth.AuthConfig, g *echo.Group) *UserHandler {
	uh := &UserHandler{
		sysOperator: sysOp,
		group:       g,
		authConf:    ac,
	}
	fmt.Println(sysOp)
	return uh
}

func (uh *UserHandler) RegisterRoutes() {
	uh.group.POST("/users/info", uh.GetById, uh.authConf.IsAuthorized)
	uh.group.GET("/users/all", uh.GetAll, uh.authConf.IsAuthorized)
	uh.group.POST("/users", uh.Save, uh.authConf.IsAuthorized)
	uh.group.PUT("/users", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/users", uh.Delete, uh.authConf.IsAuthorized)
	uh.group.PUT("/clients/update/password", uh.UpdatePassword, uh.authConf.IsAuthorized)
}

func (uh *UserHandler) Save(c echo.Context) error {
	d, l := executor.GenericPost(c, uh.sysOperator, model.User{})
	return HandleRequestResult(c, d, l)
}

func (uh *UserHandler) Update(c echo.Context) error {
	d, l := executor.GenericUpdate(c, uh.sysOperator, model.User{})
	return HandleRequestResult(c, d, l)
}

func (uh *UserHandler) Delete(c echo.Context) error {
	d, l := executor.GenericDelete(c, uh.sysOperator, model.User{})
	return HandleRequestResult(c, d, l)
}

func (uh *UserHandler) GetById(c echo.Context) error {
	d, l := executor.GenericGetById(c, uh.sysOperator, model.User{})
	return HandleRequestResult(c, d, l)
}

func (uh *UserHandler) GetAll(c echo.Context) error {
	d, l := executor.GenericGetAll(c, uh.sysOperator, []model.User{})
	return HandleRequestResult(c, d, l)
}

func (uh *UserHandler) UpdatePassword(c echo.Context) error {
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("UpdatePassword ")
	var newPassword string

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	var passwords changePasswordWrapper
	passwords, logger.Log = BindAny(c, passwords)

	if logger.Err != nil {
		return logger.Err
	}

	id := GetUIDFromContextToken(c)
	mu := model.User{ID: id}

	mu, logger.Log = executor.GenericGetWithConstraint(c, uh.sysOperator, mu, "id=?", fmt.Sprint(id))
	if logger.Err != nil {
		return logger.Err
	}

	logger.Log = auth.CheckPasswordHash(passwords.OldPassword, mu.Password)
	if logger.Err != nil {
		return logger.Err
	}

	newPassword, logger.Log = auth.GeneratehashPassword(passwords.NewPassword)
	if logger.Err != nil {
		return logger.Err
	}

	mu.Password = newPassword
	d, l := executor.GenericUpdate(c, uh.sysOperator, mu)
	return c.JSON(l.Code, d)

}
