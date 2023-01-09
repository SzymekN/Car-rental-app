package controller

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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
	uh.group.POST("/rentals/rent-for-user", uh.RentForUser, uh.authConf.IsAuthorized)
	uh.group.POST("/rentals/end", uh.EndRent, uh.authConf.IsAuthorized)
	uh.group.POST("/rentals/self", uh.SaveSelf, uh.authConf.IsAuthorized)
	uh.group.GET("/rentals/get-active", uh.GetActiveRentals, uh.authConf.IsAuthorized)
	uh.group.POST("/rentals/save-image-before", uh.SaveImageBefore, uh.authConf.IsAuthorized)
	uh.group.POST("/rentals/save-image-after", uh.SaveImageAfter, uh.authConf.IsAuthorized)
}

func (uh *RentalHandler) Save(c echo.Context) error {
	d, l := executor.GenericPost(c, uh.sysOperator, model.Rental{})
	return HandleRequestResult(c, d, l)
}

func (uh *RentalHandler) RentForUser(c echo.Context) error {

	mrw := model.RentForUserWrapper{}
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("SelfRental ")

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	mrw, logger.Log = executor.BindData(c, mrw)
	if logger.Err != nil {
		return logger.Err
	}

	var cid int
	db := uh.sysOperator.GetDB()
	// var id int
	result := db.Model(&model.User{}).Joins(" join client on user.id = client.user_id").Select("client.id").Where("email=?", mrw.Email)
	if err := result.Error; err != nil {
		logger.Log = producer.Log{
			Key:  "err",
			Msg:  "Couldn't get client id",
			Err:  err,
			Code: http.StatusInternalServerError,
		}
		return logger.Err
	}

	result.Find(&cid)
	mr := mrw.Rental
	mr.ClientID = cid

	//sprawdzanie czy fura jest dostępna

	start := mr.StartDate.Format("2006-01-02")
	end := mr.EndDate.Format("2006-01-02")

	if start > end {
		logger.Err = errors.New("Wrong dates")
		return logger.Err
	}

	result = db.Debug().Model(&model.Vehicle{}).Select("vehicle.id").Where("id not in (SELECT vehicle_id FROM `rental` where (start_date between ? and ? and end_date between ? and ?) and vehicle_id=?) and id=?", start, end, start, end, mr.VehicleID, mr.VehicleID).Scan(&model.Vehicle{})

	logger.Log = executor.CheckResultError(result)

	if logger.Log.Err != nil {
		return logger.Log.Err
	}

	logger.Log = executor.CheckIfAffected(result)

	if logger.Err != nil {
		logger.Msg = "car not available in given period"
		return logger.Err
	}

	mr, logger.Log = executor.GenericPost(c, uh.sysOperator, mr)
	if logger.Err != nil {
		return logger.Err
	}

	// mr, logger.Log = executor.GenericGetById(c, uh.sysOperator, mr)
	// if logger.Err != nil {
	// 	return logger.Err
	// }

	logger.Log.Code = http.StatusOK
	logger.Log.Key = "info"
	logger.Log.Msg = fmt.Sprintf("[INFO] completed, HTTP: %v", logger.Log.Code)
	return c.JSON(logger.Code, mr)

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

	mr, logger.Log = executor.GenericPost(c, uh.sysOperator, mr)
	if logger.Err != nil {
		return logger.Err
	}

	logger.Log.Code = http.StatusOK
	logger.Log.Key = "info"
	logger.Log.Msg = fmt.Sprintf("[INFO] completed, HTTP: %v", logger.Log.Code)
	return c.JSON(logger.Code, mr)
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

	logger.Log.Code = http.StatusOK
	logger.Log.Key = "info"
	logger.Log.Msg = fmt.Sprintf("[INFO] completed, HTTP: %v", logger.Log.Code)
	return c.JSON(l.Code, mrs)
}

func (uh *RentalHandler) EndRent(c echo.Context) error {
	mr := model.Rental{}
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("EndRental ")
	db := uh.sysOperator.GetDB()

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	mr, logger.Log = executor.BindData(c, mr)
	if logger.Err != nil && mr.GetId() < 0 {
		return logger.Err
	}

	mr.EndDate = time.Now()
	result := db.Debug().Updates(&mr)
	logger.Log = executor.CheckResultError(result)
	if logger.Log.Err != nil {
		return logger.Err
	}

	logger.Log = executor.CheckIfAffected(result)
	if logger.Log.Err != nil {
		logger.Log.Msg = "row not updated, no new values"
		return logger.Err
	}

	logger.Log.Code = http.StatusOK
	logger.Log.Key = "info"
	logger.Log.Msg = fmt.Sprintf("%s [INFO] completed, HTTP: %v", prefix, logger.Log.Code)
	return c.JSON(logger.Code, mr)
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

func (uh *RentalHandler) GetActiveRentals(c echo.Context) error {
	mr := []model.Rental{}
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("GetActiveRentals ")
	db := uh.sysOperator.GetDB()
	today := time.Now()

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	result := db.Debug().Where("end_date >= ?", today).Find(&mr)
	result.Scan(&mr)
	logger.Log = executor.CheckResultError(result)

	if logger.Log.Err != nil {
		return logger.Log.Err
	}

	logger.Log.Code = http.StatusOK
	logger.Log.Key = "info"
	logger.Log.Msg = fmt.Sprintf("[INFO] completed, HTTP: %v", logger.Log.Code)
	return c.JSON(logger.Code, mr)
}

type ImageWrapper struct {
	Id     int    `json:"rental_id"`
	Images string `json:"img"`
}

func (uh *RentalHandler) SaveImageBefore(c echo.Context) error {
	iw := ImageWrapper{}
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("SaveImage ")
	// db := uh.sysOperator.GetDB()

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	iw, logger.Log = executor.BindData(c, iw)
	if logger.Err != nil {
		return logger.Err
	}

	dir := "images/rentals/before/" + fmt.Sprint(iw.Id)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	rawDecodedText, err := base64.StdEncoding.DecodeString(iw.Images)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(dir+"/"+fmt.Sprint(time.Now().UnixNano())+".jpg", []byte(rawDecodedText), 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Println("I saved your image buddy!")
	}

	logger.Log.Code = http.StatusOK
	logger.Log.Key = "info"
	logger.Log.Msg = fmt.Sprintf("[INFO] completed, HTTP: %v", logger.Log.Code)
	return c.JSON(logger.Code, producer.GenericMessage{Message: "Images saved"})
}

func (uh *RentalHandler) SaveImageAfter(c echo.Context) error {
	iw := ImageWrapper{}
	logger := uh.sysOperator.SystemLogger
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("SaveImage ")
	// db := uh.sysOperator.GetDB()

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	iw, logger.Log = executor.BindData(c, iw)
	if logger.Err != nil {
		return logger.Err
	}

	dir := "images/rentals/after/" + fmt.Sprint(iw.Id)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	rawDecodedText, err := base64.StdEncoding.DecodeString(iw.Images)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(dir+"/"+fmt.Sprint(time.Now().UnixNano())+".jpg", []byte(rawDecodedText), 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Println("I saved your image buddy!")
	}

	logger.Log.Code = http.StatusOK
	logger.Log.Key = "info"
	logger.Log.Msg = fmt.Sprintf("[INFO] completed, HTTP: %v", logger.Log.Code)
	return c.JSON(logger.Code, producer.GenericMessage{Message: "Images saved"})
}
