package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/SzymekN/Car-rental-app/pkg/storage"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Checks for the username in the db
func GetOperator(username string) (auth.Operator, error) {
	conn := storage.MysqlConn.GetDBInstance()
	o := auth.Operator{}
	result := conn.Where(&auth.Operator{Username: username}).Find(&o)
	err := result.Error

	if err != nil {
		fmt.Println(err)
		return o, err
	}

	if result.RowsAffected < 1 {
		return o, errors.New("Operator not found")
	}

	return o, nil
}

// save signed in user to the db
func SaveOperator(op auth.Operator) error {
	pg := storage.MysqlConn.GetDBInstance()
	if err := pg.Create(&op).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Checks all passed credentials and saves user to the database
func SignUp(c echo.Context) error {

	// user to save in the database
	var op auth.Operator
	// error got while executing
	var err error
	// HTTP status code sent as a response
	var status int
	// key and message sent to kafka brokers
	k, msg := "", "userapi.operators"

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	// try saving data got in the request to the Operator datatype
	if err = c.Bind(&op); err != nil {
		status = http.StatusBadRequest
		k = op.Username
		msg += "[" + k + "] SignUp error: incorrect credentials, HTTP: " + strconv.Itoa(status)
		return err
	}

	// check if user already exists
	_, err = GetOperator(op.Username)

	k = op.Username
	if err == nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SignUp error: username in use, HTTP: " + strconv.Itoa(status)
		err = errors.New("user exists")
		return err
	}

	// hash password
	op.Password, err = auth.GeneratehashPassword(op.Password)
	if err != nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SignUp error: couldn't generate hash, HTTP: " + strconv.Itoa(status)
		return err
	}

	//insert user details to database
	err = SaveOperator(op)
	if err != nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SignUp error: insert query error, HTTP: " + strconv.Itoa(status)
		return err
	}

	status = http.StatusOK
	msg += "[" + k + "] SignUp completed: user signed up, HTTP: " + strconv.Itoa(status)
	return c.JSON(http.StatusOK, op)

}

// revokes valid jwt token. Sends the token to Redis
func SignOut(c echo.Context) error {
	var err error
	var status int = 200
	k, msg := "SignOut", "userapi.operators "

	defer func() {
		producer.ProduceMessage(k, msg)
		c.JSON(status, &model.GenericMessage{Message: msg})
	}()

	// retrieve token from the request header
	user := c.Get("user").(*jwt.Token)
	// get raw token string
	token := user.Raw
	// retrieve from string all claims
	claims := user.Claims.(jwt.MapClaims)
	// get expire date from claims
	exp := claims["exp"]
	var duration float64

	// check if token is populated and try reflecting to float
	if expFloat, ok := exp.(float64); ok && token != "" {
		duration = expFloat - float64(time.Now().Unix())
	} else {
		status = http.StatusBadRequest
		msg += "SignOut error: couldn't retrieve token, HTTP: " + strconv.Itoa(status)
		return nil
	}

	// token already not valid
	if duration < 1 {
		msg += "SignOut error: duration lesser tha 0, HTTP: " + strconv.Itoa(status)
		return nil
	}

	// save token in Redis in order to blacklist it
	err = auth.SetToken(token, time.Duration(duration))
	if err != nil {
		status = http.StatusInternalServerError
		msg += "SignOut error: couldn't write to DB, HTTP: " + strconv.Itoa(status)
		return err
	}

	fmt.Println("SIGNED  OUT")
	msg += "SignOut completed, HTTP: " + strconv.Itoa(status)
	return nil
}

// sign in a user
func SignIn(c echo.Context) error {

	var authDetails auth.Authentication
	var err error
	var status int
	k, msg := "", "userapi.operators"

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err = c.Bind(&authDetails); err != nil {
		status = http.StatusBadRequest
		k = authDetails.Username
		msg += "[" + k + "] SignIn error: incorrect credentials, HTTP: " + strconv.Itoa(status)
		return err
	}

	// check if user exists
	var authUser auth.Operator
	authUser, err = GetOperator(authDetails.Username)

	k = authDetails.Username
	if err != nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SignIn error: user doesn't exist, HTTP: " + strconv.Itoa(status)
		return err
	}

	// check if password is correct
	check := auth.CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		status = http.StatusBadRequest
		msg += "[" + k + "] SignIn error: incorrect password, HTTP: " + strconv.Itoa(status)
		err = errors.New("Incorrect password")
		return err
	}

	// generate token based on username and role
	var validToken string
	validToken, err = auth.GenerateJWT(authDetails.Username, authUser.Role)
	if err != nil {
		status = http.StatusInternalServerError
		msg += "[" + k + "] SignIn error: couldn't generate token, HTTP: " + strconv.Itoa(status)
		return err
	}

	var token auth.Token
	token.Username = authUser.Username
	token.Role = authUser.Role
	token.TokenString = validToken
	status = http.StatusOK
	msg += "[" + k + "] SignIn completed: user signed in, HTTP: " + strconv.Itoa(status)
	return c.JSON(status, token)
}
