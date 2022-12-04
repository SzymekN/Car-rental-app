package controller

import (
	"fmt"
	"net/http"

	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// type BasicControl[T model.GenericModel] struct {
// 	model T
// }

// to będzie podstawą do dziedziczenia albo i nie
// type BasicHandler interface {
// 	Post(c echo.Context) error
// 	Update(c echo.Context) error
// 	Delete(c echo.Context) error
// 	GetById(c echo.Context) error
// 	GetAll(c echo.Context) error
// }

func GenericPost[T model.GenericModel](c echo.Context, db *gorm.DB, dataModel T) error {

	log := model.Log{}
	prefix := fmt.Sprintf("POST {%T}", dataModel)

	// before exiting function send message to logs and response to user
	defer func() {
		log.Msg = fmt.Sprintf("%s %s", prefix, log.Msg)
		log.Produce(c)
	}()

	// try saving data from user request to provided model.User datatype
	dataModel, log = BindData(c, dataModel)
	if log.Err != nil {
		return log.Err
	}

	// save user in the db
	log = Insert(c, db, dataModel)
	if log.Err != nil {
		return log.Err
	}

	log.Code = http.StatusOK
	log.Key = "info"
	log.Msg = fmt.Sprintf("[INFO] insert succesful, HTTP: %v", log.Code)
	return nil

}

func GenericGetAll[T any](c echo.Context, db *gorm.DB, dataModel []T) error {

	log := model.Log{}
	prefix := fmt.Sprintf("GetALL {%T}", dataModel)

	defer func() {
		log.Msg = fmt.Sprintf("%s %s", prefix, log.Msg)
		log.Produce(c)
	}()

	result := db.Find(&dataModel)
	log = CheckResultError(result)

	if log.Err != nil {
		return log.Err
	}

	log.Code = http.StatusOK
	log.Key = "info"
	log.Msg = fmt.Sprintf("[INFO] completed, HTTP: %v", log.Code)
	return c.JSON(log.Code, dataModel)
}

// tu jeszzce spróbować dodać odbiorcę który będzie se miał bazę danych
func GenericGetById[T model.GenericModel](c echo.Context, db *gorm.DB, dataModel T) error {

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

	result := db.Find(&dataModel, id)
	log = CheckResultError(result)
	if log.Err != nil {
		return log.Err
	}

	log = CheckIfAffected(result)
	if log.Err != nil {
		return log.Err
	}

	log.Code = http.StatusOK
	log.Key = "info"
	log.Msg = fmt.Sprintf("[INFO] read completed, id: {%v} HTTP: %v", id, log.Code)
	return c.JSON(log.Code, dataModel)
}

func GenericDelete[T model.GenericModel](c echo.Context, db *gorm.DB, dataModel T) error {

	log := model.Log{}
	prefix := fmt.Sprintf("DELETE {%T}", dataModel)

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

	result := db.Delete(dataModel, id)
	log = CheckResultError(result)
	if log.Err != nil {
		return log.Err
	}

	log = CheckIfAffected(result)
	if log.Err != nil {
		return log.Err
	}

	log.Code = http.StatusOK
	log.Key = "info"
	log.Msg = fmt.Sprintf("[INFO] delete completed, id: {%v} HTTP: %v", id, log.Code)
	return nil
}

func GenericUpdate[T model.GenericModel](c echo.Context, db *gorm.DB, dataModel T) error {

	log := model.Log{}
	prefix := fmt.Sprintf("UPDATE {%T}", dataModel)
	existingModel := dataModel

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

	result := db.Debug().Find(&existingModel, id)

	log = CheckResultError(result)
	if log.Err != nil {
		prefix += " get error"
		return log.Err
	}

	log = CheckIfAffected(result)
	if log.Err != nil {
		prefix += " not found"
		return log.Err
	}

	// dataModel, log = BindData(c, dataModel)
	// if log.Err != nil {
	// 	return log.Err
	// }

	result = db.Debug().Save(&dataModel)
	log = CheckResultError(result)
	if log.Err != nil {
		return log.Err
	}

	log = CheckIfAffected(result)
	if log.Err != nil {
		return log.Err
	}

	log.Code = http.StatusOK
	log.Key = "info"
	log.Msg = fmt.Sprintf("[INFO] update completed, id: {%v} HTTP: %v", id, log.Code)
	return c.JSON(log.Code, dataModel)
}
