package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/SzymekN/Car-rental-app/pkg/storage"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BasicHandler interface {
	Post(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetById(c echo.Context) error
	GetAll(c echo.Context) error
}

func GenericPost[T any](c echo.Context, dataModel T) error {

	// error got while executing function
	var err error
	// HTTP  code to send as a response
	var status int
	// key for logger and message to save
	var k, msg string
	k = "err"
	msg = fmt.Sprintf("POST {%T}", dataModel)

	// before exiting function send message to logs and response to user
	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	// try saving data from user request to provided model.User datatype
	if err = c.Bind(&dataModel); err != nil {
		status = http.StatusBadRequest
		msg += "[ERROR]: incorrect parameters, HTTP: " + strconv.Itoa(status)
		return err
	}

	db := storage.MysqlConn.GetDBInstance()

	// save user in the db
	if err = db.Create(&dataModel).Error; err != nil {
		status = http.StatusInternalServerError
		msg += "[ERROR]: post query error, HTTP: " + strconv.Itoa(status)
		return err
	}

	status = http.StatusOK
	k = "info"
	msg += "[INFO] insert succesful, HTTP: " + strconv.Itoa(status)
	return c.JSON(status, dataModel)

}

func GenericGetAll[T any](c echo.Context, dataModel T) error {
	// error got while executing function
	var err error
	// HTTP  code to send as a response
	var status int
	// key for logger and message to save
	var k, msg string
	k = "err"
	msg = fmt.Sprintf("GetALL {%T}", dataModel)

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	db := storage.MysqlConn.GetDBInstance()
	if err := db.Find(&dataModel).Error; err != nil {
		status = http.StatusNotFound
		msg += "[ERROR] couldn't get all, HTTP: " + strconv.Itoa(status)
		return err
	}

	status = http.StatusOK
	k = "info"
	msg += "[INFO] completed, HTTP: " + strconv.Itoa(status)
	return c.JSON(status, dataModel)
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

// tu jeszzce spróbować dodać odbiorcę który będzie se miał bazę danych
func GenericGetById2[T model.GenericModel](c echo.Context, dataModel T) error {

	log := model.Log{}
	prefix := fmt.Sprintf("GetByID {%T}", dataModel)

	defer func() {
		log.Msg = fmt.Sprintf("%s %s", prefix, log.Msg)
		log.Produce(c)
	}()

	dataModel, log = BindData(c, dataModel)
	if log.Err != nil {
		return log.Err
	}

	id := dataModel.GetId()

	log = CheckID(id)
	if log.Err != nil {
		return log.Err
	}

	db := storage.MysqlConn.GetDBInstance()
	log = CheckIfAffected(db.Find(&dataModel, id))
	if log.Err != nil {
		return log.Err
	}

	log.Code = http.StatusOK
	log.Key = "info"
	log.Msg = fmt.Sprintf("[INFO] completed: user read, id: {%v} HTTP: %v", id, log.Code)
	return nil
}

func GenericGetById[T model.GenericModel](c echo.Context, dataModel T) error {
	// error got while executing function
	var err error
	// HTTP  code to send as a response
	var status int
	// key for logger and message to save
	var k, msg string
	k = "err"
	msg = fmt.Sprintf("GetByID {%T}", dataModel)

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err = c.Bind(&dataModel); err != nil {
		status = http.StatusBadRequest
		msg += fmt.Sprintf("[ERROR]: couldn't get id from request, HTTP: %v", status)
		return err
	}

	id := dataModel.GetId()
	if id < 1 {
		status = http.StatusBadRequest
		msg += fmt.Sprintf("[ERROR] invalid id: {%v}, HTTP: %v", id, status)
		err = fmt.Errorf("invalid id: {%v}", id)
		return err
	}

	db := storage.MysqlConn.GetDBInstance()
	result := db.Find(&dataModel, id)

	if result.RowsAffected < 1 {
		status = http.StatusNotFound
		msg += fmt.Sprintf("[ERROR]: couldn't get user, id: {%v}, HTTP: %v", id, status)
		err = fmt.Errorf("couldn't get %T", dataModel)
		return err
	}

	status = http.StatusOK
	k = "info"
	msg += "[INFO] completed: user read, id: {" + strconv.Itoa(id) + "} HTTP: " + strconv.Itoa(status)
	return c.JSON(status, dataModel)
}

func DeleteUser[T model.GenericModel](c echo.Context, dataModel T) error {
	// error got while executing function
	var err error
	// HTTP  code to send as a response
	var status int
	// key for logger and message to save
	var k, msg string
	k = "err"
	msg = fmt.Sprintf("Delete {%T}", dataModel)

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err = c.Bind(&dataModel); err != nil {
		status = http.StatusBadRequest
		msg += fmt.Sprintf("[ERROR]: couldn't get id from request, HTTP: %v", status)
		return err
	}

	id := dataModel.GetId()
	if id < 1 {
		k = "unknown"
		status = http.StatusBadRequest
		msg += "[ERROR] invalid id: {" + strconv.Itoa(id) + "} HTTP: " + strconv.Itoa(status)
		return err
	}

	db := storage.MysqlConn.GetDBInstance()
	k = strconv.Itoa(id)
	result := db.Delete(dataModel, id)

	if result.RowsAffected < 1 {
		status = http.StatusNotFound
		msg += "[" + k + "] DeleteUser error: user doesn't exist, HTTP: " + strconv.Itoa(status)
		err = errors.New("user doesn't exist")
		return err
	}

	status = http.StatusOK
	msg += "[" + k + "] DeleteUser completed: user deleted, HTTP: " + strconv.Itoa(status)
	return c.JSON(status, &model.GenericMessage{Message: msg})
}
