package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"

	"github.com/labstack/echo/v4"
)

// swagger:route POST /api/v1/users/save users_v1 postUserV1
// Save user to  database.
//
//		Consumes:
//	   - application/json
//	 Produces:
//	   - application/json
//
// responses:
//
//	200: userResponse
//	500: errorResponse

type UsersController struct {
	MainController
}

type UsersHandler interface {
	SaveUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	GetUserById(c echo.Context) error
	GetUsers(c echo.Context) error
}

func (uc *UsersController) SaveUser(c echo.Context) error {
	return GenericPost(c, model.User{})
}

func (uc *UsersController) UpdateUser(c echo.Context) error {
	// HTTP status code send as a response
	var status int
	// id of a user to update
	var id int
	// error
	var err error
	id, err = strconv.Atoi(c.Param("id"))
	k, msg := "", "userapi_v1.users"

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()
	if err != nil {
		k = "unknown"
		status = http.StatusBadRequest
		msg += "[" + k + "] UpdateUser error: incorrect id, HTTP: " + strconv.Itoa(status)
		return err
	}

	k = strconv.Itoa(id)
	db := uc.GetDB()
	user := model.User{}
	result := db.Find(&user, id)

	if result.RowsAffected < 1 {
		status = http.StatusNotFound
		msg += "[" + k + "] UpdateUser error: user doesn't exist, HTTP: " + strconv.Itoa(status)
		err = errors.New("user doesn't exist")
		return err
	}

	if err = c.Bind(&user); err != nil {
		status = http.StatusBadRequest
		msg += "[" + k + "] UpdateUser error: incorrect parameters, HTTP: " + strconv.Itoa(status)
		return err
	}

	user.ID = id
	result = db.Save(&user)
	if result.RowsAffected < 1 {
		status = http.StatusInternalServerError
		msg += "[" + k + "] UpdateUser error: update query error, HTTP: " + strconv.Itoa(status)
		err = errors.New("update query error")
		return err
	}

	status = http.StatusOK
	msg += "[" + k + "] UpdateUser completed: user updated, HTTP: " + strconv.Itoa(status)
	return c.JSON(status, user)
}

func (uc *UsersController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	var status int
	k, msg := "", "userapi_v1.users"

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err != nil {
		k = "unknown"
		status = http.StatusBadRequest
		msg += "[" + k + "] DeleteUser error: incorrect id, HTTP: " + strconv.Itoa(status)
		return err
	}

	k = strconv.Itoa(id)
	db := uc.GetDB()
	result := db.Delete(&model.User{}, id)

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

func (uc *UsersController) GetUserById(c echo.Context) error {
	return GenericGetById2(c, model.User{})
}

func (uc *UsersController) GetUsers(c echo.Context) error {
	return GenericGetAll(c, model.User{})
}
