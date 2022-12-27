package controller

import (
	"errors"
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
	uh.group.GET("/rentals/self", uh.GetSelf, uh.authConf.IsAuthorized)
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

	var cid int
	cid, logger.Log = GetCIDFromContextToken(c, uh.sysOperator)
	if logger.Err != nil {
		return logger.Err
	}
	mr.ClientID = cid

	//sprawdzanie czy fura jest dostępna

	start := mr.StartDate.Format("2006-01-02")
	end := mr.EndDate.Format("2006-01-02")

	if start > end {
		logger.Err = errors.New("Wrong dates")
		return logger.Err
	}

	db := uh.sysOperator.GetDB()
	result := db.Debug().Model(&model.Vehicle{}).Select("vehicle.id").Where("id not in (SELECT vehicle_id FROM `rental` where (start_date between ? and ? and end_date between ? and ?) and vehicle_id=?) and id=?", start, end, start, end, mr.VehicleID, mr.VehicleID).Scan(&model.Vehicle{})

	logger.Log = executor.CheckResultError(result)

	if logger.Log.Err != nil {
		return logger.Log.Err
	}

	logger.Log = executor.CheckIfAffected(result)

	if logger.Err != nil {
		logger.Msg = "car not available in given period"
		return logger.Err
	}

	d, l := executor.GenericPost(c, uh.sysOperator, mr)
	if logger.Err != nil {
		return logger.Err
	}
	return c.JSON(l.Code, d)
}

func (uh *RentalHandler) GetSelf(c echo.Context) error {
	mrs := []model.Rental{}
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("SelfRental ")

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	var cid int
	cid, logger.Log = GetCIDFromContextToken(c, uh.sysOperator)
	if logger.Err != nil {
		return logger.Err
	}

	mrs, l := executor.GenericGetAllWithConstraint(c, uh.sysOperator, mrs, "client_id = ?", fmt.Sprint(cid))

	if logger.Err != nil {
		return logger.Err
	}

	return c.JSON(l.Code, mrs)
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
