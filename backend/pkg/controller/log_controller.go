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

type LogHandler struct {
	sysOperator producer.SystemOperator
	authConf    auth.AuthConfig
	group       *echo.Group
}

func NewLogHandler(sysOp producer.SystemOperator, ac auth.AuthConfig, g *echo.Group) *LogHandler {
	uh := &LogHandler{
		sysOperator: sysOp,
		group:       g,
		authConf:    ac,
	}
	fmt.Println(sysOp)
	return uh
}

func (uh *LogHandler) RegisterRoutes() {
	uh.group.POST("/logs/info", uh.GetById, uh.authConf.IsAuthorized)
	uh.group.GET("/logs/all", uh.GetAll, uh.authConf.IsAuthorized)
	uh.group.POST("/logs", uh.Save, uh.authConf.IsAuthorized)
	uh.group.POST("/logs/get-range", uh.GetFromRange, uh.authConf.IsAuthorized)
	uh.group.PUT("/logs", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/logs", uh.Delete, uh.authConf.IsAuthorized)
}

func (uh *LogHandler) Save(c echo.Context) error {
	d, l := executor.GenericPost(c, uh.sysOperator, model.Notification{})
	return HandleRequestResult(c, d, l)
}

func (uh *LogHandler) Update(c echo.Context) error {
	d, l := executor.GenericUpdate(c, uh.sysOperator, model.Notification{})
	return HandleRequestResult(c, d, l)
}

func (uh *LogHandler) Delete(c echo.Context) error {
	d, l := executor.GenericDelete(c, uh.sysOperator, model.Notification{})
	return HandleRequestResult(c, d, l)
}

func (uh *LogHandler) GetById(c echo.Context) error {
	d, l := executor.GenericGetById(c, uh.sysOperator, model.Notification{})
	return HandleRequestResult(c, d, l)
}

func (uh *LogHandler) GetAll(c echo.Context) error {
	d, l := executor.GenericGetAll(c, uh.sysOperator, []model.Log{})
	return HandleRequestResult(c, d, l)
}

type LogRequest struct {
	Batch_nr int `json:"batch_nr"`
}

func (uh *LogHandler) GetFromRange(c echo.Context) error {

	db := uh.sysOperator.GetDB()
	prefix := fmt.Sprintf("GetLogFromRange ")
	logs := []model.Log{}
	lReq := LogRequest{}
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	lReq, logger.Log = executor.BindData(c, lReq)
	if logger.Err != nil {
		return logger.Err
	}

	batchNr := lReq.Batch_nr

	result := db.Debug().Model(&model.Log{}).Select("*").Where("id < ((SELECT MAX(id) from log)-50*?)", batchNr).Order("id desc").Limit(50)
	result.Scan(&logs)
	logger.Log = executor.CheckResultError(result)

	if logger.Log.Err != nil {
		return logger.Log.Err
	}

	logger.Log = executor.CheckIfAffected(result)

	if logger.Log.Err != nil {
		logger.Log = producer.Log{
			Key:  "info",
			Msg:  fmt.Sprintf("[INFO] completed, HTTP: %v", http.StatusNotFound),
			Code: http.StatusNotFound,
			Err:  nil}
		fmt.Println(logger.Log)
		return c.JSON(404, producer.GenericMessage{Message: "No more logs"})
	}

	logger.Log.Code = http.StatusOK
	logger.Log.Key = "info"
	logger.Log.Msg = fmt.Sprintf("[INFO] completed, HTTP: %v", logger.Log.Code)
	return c.JSON(200, logs)
}
