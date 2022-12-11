package executor

import (
	"fmt"
	"net/http"

	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// type BasicOperationsInterface interface {
// 	CheckResultError(result *gorm.DB) producer.Log
// 	CheckIfAffected(result *gorm.DB) producer.Log
// 	CheckID(id int) producer.Log
// }

func Insert[T model.GenericModel](c echo.Context, db *gorm.DB, d T) producer.Log {

	log := producer.Log{}
	if err := db.Create(&d).Error; err != nil {
		log.Code = http.StatusInternalServerError
		log.Msg = fmt.Sprintf("[ERROR]: post query error, HTTP: %v", log.Code)
		log.Err = err
		log.Key = "err"
		return log
	}

	return log
}

func CheckResultError(result *gorm.DB) producer.Log {
	log := producer.Log{}
	if err := result.Error; err != nil {
		log.Key = "err"
		log.Code = http.StatusNotFound
		log.Msg = fmt.Sprintf("[ERROR] couldn't get all, HTTP: %v", log.Code)
		log.Err = err
		return log
	}
	return log
}

func BindData[T model.GenericModel](c echo.Context, d T) (T, producer.Log) {

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

func CheckID(id int) producer.Log {

	if id < 1 {
		status := http.StatusBadRequest
		msg := fmt.Sprintf("[ERROR] invalid id: {%v}, HTTP: %v", id, status)
		err := fmt.Errorf("invalid id: {%v}", id)
		log := producer.Log{
			Key:  "err",
			Msg:  msg,
			Code: status,
			Err:  err}
		return log
	}

	return producer.Log{}
}

func CheckIfAffected(result *gorm.DB) producer.Log {

	if result.RowsAffected < 1 {
		status := http.StatusBadRequest
		msg := fmt.Sprintf("[ERROR]: not found / no rows affected, HTTP: %v", status)
		err := fmt.Errorf("no rows affected")
		log := producer.Log{
			Key:  "err",
			Msg:  msg,
			Code: status,
			Err:  err}
		return log
	}
	return producer.Log{}

}
