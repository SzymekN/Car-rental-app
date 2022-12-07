package executor

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/labstack/echo/v4"
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

func GenericPost[T model.GenericModel](c echo.Context, so producer.SystemOperator, dataModel T) error {

	so.Log = producer.Log{}
	db := so.GetDB()
	prefix := fmt.Sprintf("POST {%T}", dataModel)

	// before exiting function send message to logs and response to user
	defer func() {
		so.Log.Msg = fmt.Sprintf("%s %s", prefix, so.Log.Msg)
		so.Produce(c)
	}()

	// try saving data from user request to provided model.User datatype
	dataModel, so.Log = BindData(c, dataModel)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	// save user in the db
	so.Log = Insert(c, db, dataModel)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	so.Log.Code = http.StatusOK
	so.Log.Key = "info"
	so.Log.Msg = fmt.Sprintf("[INFO] insert succesful, HTTP: %v", so.Log.Code)
	fmt.Println(so.Log)
	return nil

}

func GenericGetAll[T any](c echo.Context, so producer.SystemOperator, dataModel []T) error {

	so.Log = producer.Log{}
	db := so.GetDB()
	prefix := fmt.Sprintf("GetALL {%T}", dataModel)

	defer func() {
		so.Log.Msg = fmt.Sprintf("%s %s", prefix, so.Log.Msg)
		so.Produce(c)
	}()

	result := db.Find(&dataModel)
	so.Log = CheckResultError(result)

	if so.Log.Err != nil {
		return so.Log.Err
	}

	so.Log.Code = http.StatusOK
	so.Log.Key = "info"
	so.Log.Msg = fmt.Sprintf("[INFO] completed, HTTP: %v", so.Log.Code)
	return c.JSON(so.Log.Code, dataModel)
}

// tu jeszzce spróbować dodać odbiorcę który będzie se miał bazę danych
func GenericGetById[T model.GenericModel](c echo.Context, so producer.SystemOperator, dataModel T) error {

	so.Log = producer.Log{}
	db := so.GetDB()
	prefix := fmt.Sprintf("GetByID {%T}", dataModel)

	defer func() {
		so.Log.Msg = fmt.Sprintf("%s %s", prefix, so.Log.Msg)
		so.Produce(c)
	}()

	dataModel, so.Log = BindData(c, dataModel)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	id := dataModel.GetId()

	so.Log = CheckID(id)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	result := db.Debug().Find(&dataModel, id)
	so.Log = CheckResultError(result)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	so.Log = CheckIfAffected(result)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	so.Log.Code = http.StatusOK
	so.Log.Key = "info"
	so.Log.Msg = fmt.Sprintf("[INFO] read completed, id: {%v} HTTP: %v", id, so.Log.Code)
	return c.JSON(so.Log.Code, dataModel)
}

func GenericDelete[T model.GenericModel](c echo.Context, so producer.SystemOperator, dataModel T) error {

	so.Log = producer.Log{}
	db := so.GetDB()
	prefix := fmt.Sprintf("DELETE {%T}", dataModel)

	defer func() {
		so.Log.Msg = fmt.Sprintf("%s %s", prefix, so.Log.Msg)
		so.Produce(c)
	}()

	dataModel, so.Log = BindData(c, dataModel)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	id := dataModel.GetId()

	so.Log = CheckID(id)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	result := db.Delete(dataModel, id)
	so.Log = CheckResultError(result)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	so.Log = CheckIfAffected(result)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	so.Log.Code = http.StatusOK
	so.Log.Key = "info"
	so.Log.Msg = fmt.Sprintf("[INFO] delete completed, id: {%v} HTTP: %v", id, so.Log.Code)
	return nil
}

func GenericUpdate[T model.GenericModel](c echo.Context, so producer.SystemOperator, dataModel T) error {

	so.Log = producer.Log{}
	db := so.GetDB()
	prefix := fmt.Sprintf("UPDATE {%T}", dataModel)
	existingModel := dataModel

	defer func() {
		so.Log.Msg = fmt.Sprintf("%s %s", prefix, so.Log.Msg)
		so.Produce(c)
	}()

	dataModel, so.Log = BindData(c, dataModel)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	id := dataModel.GetId()

	so.Log = CheckID(id)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	result := db.Debug().Find(&existingModel, id)

	so.Log = CheckResultError(result)
	if so.Log.Err != nil {
		prefix += " get error"
		return so.Log.Err
	}

	so.Log = CheckIfAffected(result)
	if so.Log.Err != nil {
		prefix += " not found"
		return so.Log.Err
	}

	// dataModel, so.Log = BindData(c, dataModel)
	// if so.Log.Err != nil {
	// 	return so.Log.Err
	// }

	result = db.Debug().Updates(&dataModel)
	so.Log = CheckResultError(result)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	so.Log = CheckIfAffected(result)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	so.Log.Code = http.StatusOK
	so.Log.Key = "info"
	so.Log.Msg = fmt.Sprintf("[INFO] update completed, id: {%v} HTTP: %v", id, so.Log.Code)
	return c.JSON(so.Log.Code, dataModel)
}

// tu jeszzce spróbować dodać odbiorcę który będzie se miał bazę danych
func GenericGetWithConstraint[T model.GenericModel](c echo.Context, so producer.SystemOperator, dataModel T, constraint string, values ...string) error {

	so.Log = producer.Log{}
	db := so.GetDB()
	prefix := fmt.Sprintf("GetByID {%T}", dataModel)

	defer func() {
		so.Log.Msg = fmt.Sprintf("%s %s", prefix, so.Log.Msg)
		so.Produce(c)
	}()

	dataModel, so.Log = BindData(c, dataModel)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	// id := dataModel.GetId()

	// so.Log = CheckID(id)
	// if so.Log.Err != nil {
	// 	return so.Log.Err
	// }

	result := db.Debug().Where(constraint, values).Find(&dataModel)
	so.Log = CheckResultError(result)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	so.Log = CheckIfAffected(result)
	if so.Log.Err != nil {
		return so.Log.Err
	}

	so.Log.Code = http.StatusOK
	so.Log.Key = "info"
	so.Log.Msg = fmt.Sprintf("[INFO] read completed, constraint: {%s}, values: {"+strings.Join(values, ", ")+"} HTTP: %v", constraint, so.Log.Code)
	return c.JSON(so.Log.Code, dataModel)
}
