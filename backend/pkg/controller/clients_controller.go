package controller

import (
	"fmt"
	"net/http"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/executor"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
)

type ClientHandler struct {
	sysOperator producer.SystemOperator
	authConf    auth.AuthConfig
	group       *echo.Group
}

func NewClientHandler(sysOp producer.SystemOperator, ac auth.AuthConfig, g *echo.Group) *ClientHandler {
	uh := &ClientHandler{
		sysOperator: sysOp,
		group:       g,
		authConf:    ac,
	}
	fmt.Println(sysOp)
	return uh
}

func (uh *ClientHandler) RegisterRoutes() {
	uh.group.GET("/clients", uh.GetById, uh.authConf.IsAuthorized)
	uh.group.GET("/clients/all", uh.GetAll, uh.authConf.IsAuthorized)
	uh.group.PUT("/clients", uh.Update, uh.authConf.IsAuthorized)
	uh.group.PUT("/clients/self", uh.UpdateSelf, uh.authConf.IsAuthorized)
	uh.group.POST("/clients", uh.Save, uh.authConf.IsAuthorized)
	uh.group.DELETE("/clients", uh.Delete, uh.authConf.IsAuthorized)
	uh.group.DELETE("/clients/self", uh.DeleteSelf, uh.authConf.IsAuthorized)
}

func GetClientID(c echo.Context, so producer.SystemOperator, uid int) (int, producer.Log) {
	db := so.GetDB()
	var id int
	result := db.Model(&model.Client{}).Select("ID").Where("user_id=?", uid)

	if err := result.Error; err != nil {
		log := producer.Log{
			Key:  "err",
			Msg:  "Couldn't get client id",
			Err:  err,
			Code: http.StatusInternalServerError,
		}
		return -1, log
	}

	result.Find(&id)
	return id, producer.Log{}

}

// TODO error check
func GetUIDFromContextToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	uid := int(claims["id"].(float64))
	return uid
}

func GetCIDFromContextToken(c echo.Context, so producer.SystemOperator) (int, producer.Log) {

	log := producer.Log{}
	var uid, cid int

	uid = GetUIDFromContextToken(c)

	// if log.Err != nil {
	// 	return -1, log
	// }

	cid, log = GetClientID(c, so, uid)
	if log.Err != nil {
		return -1, log
	}

	return cid, log
}

type passwordWrapper struct {
	Password string `json:"password"`
}

type changePasswordWrapper struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func BindPassword(c echo.Context, d passwordWrapper) (passwordWrapper, producer.Log) {

	if err := c.Bind(&d); err != nil {
		status := http.StatusBadRequest
		msg := fmt.Sprintf("[ERROR]: couldn't bind data from request, HTTP: %v", status)
		log := producer.Log{
			Key:  "err",
			Msg:  msg,
			Code: status,
			Err:  err}
		return d, log
	}

	return d, producer.Log{}
}

func BindChangePassword(c echo.Context, d changePasswordWrapper) (changePasswordWrapper, producer.Log) {

	if err := c.Bind(&d); err != nil {
		status := http.StatusBadRequest
		msg := fmt.Sprintf("[ERROR]: couldn't bind data from request, HTTP: %v", status)
		log := producer.Log{
			Key:  "err",
			Msg:  msg,
			Code: status,
			Err:  err}
		return d, log
	}

	return d, producer.Log{}
}

func (uh *ClientHandler) Save(c echo.Context) error {
	d, l := executor.GenericPost(c, uh.sysOperator, model.Client{})
	return HandleRequestResult(c, d, l)
}

func (uh *ClientHandler) Update(c echo.Context) error {
	d, l := executor.GenericUpdate(c, uh.sysOperator, model.Client{})
	return HandleRequestResult(c, d, l)
}

func (uh *ClientHandler) Delete(c echo.Context) error {
	d, l := executor.GenericDelete(c, uh.sysOperator, model.Client{})
	return HandleRequestResult(c, d, l)
}

func (uh *ClientHandler) GetById(c echo.Context) error {
	d, l := executor.GenericGetById(c, uh.sysOperator, model.Client{})
	return HandleRequestResult(c, d, l)
}

func (uh *ClientHandler) GetAll(c echo.Context) error {
	d, l := executor.GenericGetAll(c, uh.sysOperator, []model.Client{})
	return HandleRequestResult(c, d, l)
}

func (uh *ClientHandler) UpdateSelf(c echo.Context) error {
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("UpdateSelf ")

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	mc := model.Client{}
	mc, logger.Log = executor.BindData(c, mc)
	fmt.Println(mc)
	if logger.Log.Err != nil {
		return logger.Log.Err
	}

	id := GetUIDFromContextToken(c)
	mc.User.ID = id
	mc.UserID = id
	mc.ID = 0

	newmc := model.Client{}
	newmc, logger.Log = executor.GenericGetWithConstraint(c, uh.sysOperator, mc, "user_id=?", fmt.Sprint(id))
	mc.ID = newmc.ID

	if logger.Err != nil {
		return logger.Err
	}

	mc, logger.Log = executor.GenericUpdate(c, uh.sysOperator, mc)
	mc.User, logger.Log = executor.GenericUpdate(c, uh.sysOperator, mc.User)

	fmt.Println(mc)
	if logger.Err != nil {
		return logger.Err
	}

	return c.JSON(logger.Code, mc)

}

// prawdopodobnie do usunięcia
// func GetCIDFromContextToken(c echo.Context, so producer.SystemOperator) (int, producer.Log) {
// 	id := getUIDFromContextToken(c)
// 	mu := model.Client{UserID: id}
// 	mu, so.Log = executor.GenericGetWithConstraint(c, so, mu, "user_id=?", fmt.Sprint(id))
// 	return mu.ID, so.Log
// }

func (uh *ClientHandler) DeleteSelf(c echo.Context) error {
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("DeleteSelf ")

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	var pwd passwordWrapper
	pwd, logger.Log = BindPassword(c, pwd)

	if logger.Err != nil {
		return logger.Err
	}

	id := GetUIDFromContextToken(c)
	mu := model.User{ID: id}

	mu, logger.Log = executor.GenericGetWithConstraint(c, uh.sysOperator, mu, "id=?", fmt.Sprint(id))
	if logger.Err != nil {
		return logger.Err
	}

	logger.Log = auth.CheckPasswordHash(pwd.Password, mu.Password)
	if logger.Err != nil {
		return logger.Err
	}

	d, l := executor.GenericDelete(c, uh.sysOperator, mu)
	return c.JSON(l.Code, d)

}
