package controller

import (
	"fmt"
	"net/http"

	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Insert[T model.GenericModel](c echo.Context, db *gorm.DB, d T) model.Log {

	log := model.Log{}
	if err := db.Create(&d).Error; err != nil {
		log.Code = http.StatusInternalServerError
		log.Msg = fmt.Sprintf("[ERROR]: post query error, HTTP: %v", log.Code)
		log.Err = err
		log.Key = "err"
		return log
	}

	return log
}

func CheckResultError(result *gorm.DB) model.Log {
	log := model.Log{}
	if err := result.Error; err != nil {
		log.Key = "err"
		log.Code = http.StatusNotFound
		log.Msg = fmt.Sprintf("[ERROR] couldn't get all, HTTP: %v", log.Code)
		log.Err = err
		return log
	}
	return log
}

func BindData[T model.GenericModel](c echo.Context, d T) (T, model.Log) {

	if err := c.Bind(&d); err != nil {
		status := http.StatusBadRequest
		msg := fmt.Sprintf("[ERROR]: couldn't get id from request, HTTP: %v", status)
		log := model.Log{
			Key:  "err",
			Msg:  msg,
			Code: status,
			Err:  err}
		return d, log
	}

	return d, model.Log{}
}

func CheckID(id int) model.Log {

	if id < 1 {
		status := http.StatusBadRequest
		msg := fmt.Sprintf("[ERROR] invalid id: {%v}, HTTP: %v", id, status)
		err := fmt.Errorf("invalid id: {%v}", id)
		log := model.Log{
			Key:  "err",
			Msg:  msg,
			Code: status,
			Err:  err}
		return log
	}

	return model.Log{}
}

func CheckIfAffected(result *gorm.DB) model.Log {

	if result.RowsAffected < 1 {
		status := http.StatusBadRequest
		msg := fmt.Sprintf("[ERROR]: no rows affected, HTTP: %v", status)
		err := fmt.Errorf("no rows affected")
		log := model.Log{
			Key:  "err",
			Msg:  msg,
			Code: status,
			Err:  err}
		return log
	}
	return model.Log{}

}
