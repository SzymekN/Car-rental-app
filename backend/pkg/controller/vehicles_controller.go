package controller

import (
	"fmt"
	"strings"
	"time"

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
	uh.group.GET("/vehicles/available", uh.GetAvailable)
	uh.group.POST("/vehicles", uh.Save, uh.authConf.IsAuthorized)
	uh.group.PUT("/vehicles", uh.Update, uh.authConf.IsAuthorized)
	uh.group.DELETE("/vehicles", uh.Delete, uh.authConf.IsAuthorized)
}

func (uh *VehicleHandler) Save(c echo.Context) error {
	d, l := executor.GenericPost(c, uh.sysOperator, model.Vehicle{})
	return HandleRequestResult(c, d, l)
}

func (uh *VehicleHandler) Update(c echo.Context) error {
	d, l := executor.GenericUpdate(c, uh.sysOperator, model.Vehicle{})
	return HandleRequestResult(c, d, l)
}

func (uh *VehicleHandler) Delete(c echo.Context) error {
	d, l := executor.GenericDelete(c, uh.sysOperator, model.Vehicle{})
	return HandleRequestResult(c, d, l)
}

func (uh *VehicleHandler) GetById(c echo.Context) error {
	d, l := executor.GenericGetById(c, uh.sysOperator, model.Vehicle{})
	return HandleRequestResult(c, d, l)
}

func (uh *VehicleHandler) GetAll(c echo.Context) error {
	d, l := executor.GenericGetAll(c, uh.sysOperator, []model.Vehicle{})
	return HandleRequestResult(c, d, l)
}

type CustomTime struct {
	time.Time
}

type DateRange struct {
	StartDate CustomTime `json:"start_date,omitempty"`
	EndDate   CustomTime `json:"end_date,omitempty"`
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	date := t.Time.Format("2006-01-02")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}

func (uh *VehicleHandler) GetAvailable(c echo.Context) error {
	dr := DateRange{}
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("GetAvailableVehicles ")

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	dr, logger.Log = executor.BindData(c, dr)
	if logger.Err != nil && logger.Err.Error() == "EOF" {
		dr.StartDate.Time = time.Now()
		dr.EndDate.Time = time.Now()
	} else if logger.Err != nil {
		return logger.Err
	}

	t := time.Time{}
	if dr.StartDate.Time == t {
		dr.StartDate.Time = time.Now()
	}
	if dr.EndDate.Time == t {
		dr.EndDate.Time = time.Now()
	}

	start := dr.StartDate.Format(time.RFC3339)
	end := dr.EndDate.Format(time.RFC3339)

	// d, l := executor.GenericGetAllWithConstraint(c, uh.sysOperator, []model.Rental{}, "start_date not between ? and ? and end_date not between ? and ?", start, end, start, end)
	db := uh.sysOperator.DB
	vehicles := []model.Vehicle{}
	result := db.Debug().Model(&model.Vehicle{}).Select("*").Joins("join rental on vehicle.ID = rental.vehicle_id").Where("start_date not between ? and ? and end_date not between ? and ?", start, end, start, end)
	result.Scan(&vehicles)

	logger.Log = executor.CheckResultError(result)

	if logger.Log.Err != nil {
		return logger.Log.Err
	}

	return c.JSON(logger.Code, vehicles)
	// return HandleRequestResult(c, vehicles, l)
}
