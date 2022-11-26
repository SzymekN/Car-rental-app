package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/SzymekN/Car-rental-app/pkg/storage"

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
func SaveUser(c echo.Context) error {

	// user to be saved in the db
	var u model.User
	// error got while executing function
	var err error
	// HTTP  code to send as a response
	var status int
	// key for logger and message to save
	k, msg := "", "userapi_v1.users"

	// before exiting function send message to logs and response to user
	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	// try saving data from user request to provided model.User datatype
	if err = c.Bind(&u); err != nil {
		status = http.StatusBadRequest
		msg += "[" + k + "] SaveUser error: incorrect parameters, HTTP: " + strconv.Itoa(status)
		return err
	}

	k = strconv.Itoa(u.Id)
	db := storage.MysqlConn.GetDBInstance()

	// save user in the db
	if err = db.Create(&u).Error; err != nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SaveUser error: post query error, HTTP: " + strconv.Itoa(status)
		return err
	}

	status = http.StatusOK
	msg += "[" + k + "] SaveUser completed: user added, HTTP: " + strconv.Itoa(status)
	return c.JSON(status, u)
}

// swagger:route PUT /api/v1/user/{id} users_v1 putUserV1
// Updates user in  database.
//
//		Consumes:
//	   - application/json
//	 Produces:
//	   - application/json
//
// responses:
//
//	200: userResponse
//	400: errorResponse
//	404: errorResponse
//	500: errorResponse
func UpdateUser(c echo.Context) error {
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
	db := storage.MysqlConn.GetDBInstance()
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

	user.Id = id
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

// swagger:route DELETE /api/v1/user/{id} users_v1 deleteUserV1
// deletes user from  database.
//
//	Produces:
//	  - application/json
//
// responses:
//
//	200: messageResponse
//	400: errorResponse
//	404: errorResponse
func DeleteUser(c echo.Context) error {
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
	db := storage.MysqlConn.GetDBInstance()
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

// swagger:route GET /api/v1/user/{id} users_v1 getUserV1
// Gets user from  database.
//
//	Produces:
//	  - application/json
//
// responses:
//
//	200: userResponse
//	400: errorResponse
//	404: errorResponse
func GetUserById(c echo.Context) error {
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
		msg += "[" + k + "] GetUserById error: incorrect id, HTTP: " + strconv.Itoa(status)
		return err
	}

	k = strconv.Itoa(id)
	db := storage.MysqlConn.GetDBInstance()
	user := model.User{}
	result := db.Find(&user, id)

	if result.RowsAffected < 1 {
		status = http.StatusNotFound
		msg += "[" + k + "] GetUserById error: couldn't get user, HTTP: " + strconv.Itoa(status)
		err = errors.New("couldn't get user")
		return err
	}

	status = http.StatusOK
	msg += "[" + k + "] GetUserById completed: user read, HTTP: " + strconv.Itoa(status)
	return c.JSON(status, user)
}

// swagger:route GET /api/v1/users users_v1 listUsersV1
// Gets user from  database.
//
//	Produces:
//	  - application/json
//
// responses:
//
//	200: usersResponse
//	500: errorResponse
func GetUsers(c echo.Context) error {
	db := storage.MysqlConn.GetDBInstance()
	users := []model.User{}

	k, msg := "all", "userapi_v1.users"
	var status int

	defer func() {
		producer.ProduceMessage(k, msg)
	}()

	if err := db.Find(&users).Error; err != nil {
		status = http.StatusNotFound
		msg += "[" + k + "] GetUsers error: couldn't get users, HTTP: " + strconv.Itoa(status)
		return err
	}

	status = http.StatusOK
	msg += "[" + k + "] GetUsers completed: users read, HTTP: " + strconv.Itoa(status)
	return c.JSON(status, users)
}
