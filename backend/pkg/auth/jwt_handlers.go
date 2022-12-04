package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type JWTHandler struct {
	JwtC JWTControl
}

func (j JWTHandler) GetMysqlDB() *gorm.DB {
	return j.JwtC.JwtQE.Svr.GetMysqlDB()
}
func (j JWTHandler) GetSigningKey() string {
	return j.JwtC.SecretKey
}

// Checks for the username in the db
func (j JWTHandler) GetUser(email string) (model.User, error) {
	db := j.GetMysqlDB()
	u := model.User{}
	result := db.Debug().Where(&model.User{Email: email}).Find(&u)
	err := result.Error

	if err != nil {
		fmt.Println(err)
		return u, err
	}

	if result.RowsAffected < 1 {
		return u, errors.New("User not found")
	}

	return u, nil
}

// save signed in user to the db
func (j JWTHandler) SignUser(mc model.Client) error {
	db := j.GetMysqlDB()

	fmt.Println(db.Model(&model.Client{}).Association("User").Error)
	if err := db.Model(&model.Client{}).Preload("User").Debug().Create(&mc).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// Checks all passed credentials and saves user to the database
func (j JWTHandler) SignUp(c echo.Context) error {

	// user to save in the database
	var mc model.Client
	// error got while executing
	var err error
	// HTTP status code sent as a response
	var status int
	// key and message sent to kafka brokers
	k, msg := "err", "[ERROR]"

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	// try saving data got in the request to the User datatype
	if err = c.Bind(&mc); err != nil {
		status = http.StatusBadRequest
		k = mc.User.Email
		msg += " SignUp error: incorrect credentials, email: {" + k + "}, HTTP: " + strconv.Itoa(status)
		return err
	}

	// check if user already exists
	_, err = j.GetUser(mc.User.Email)

	k = mc.User.Email
	if err == nil {
		status = http.StatusInternalServerError
		msg += " SignUp error: email in use, email: {" + k + "}, HTTP: " + strconv.Itoa(status)
		err = errors.New("user exists")
		return err
	}

	// hash password
	mc.User.Password, err = j.JwtC.GeneratehashPassword(mc.User.Password)
	if err != nil {
		status = http.StatusInternalServerError
		msg += " SignUp error: couldn't generate hash, email: {" + k + "}, HTTP: " + strconv.Itoa(status)
		return err
	}

	//insert user details to database
	err = j.SignUser(mc)
	if err != nil {
		status = http.StatusInternalServerError
		msg += " SignUp error: insert query error, email: {" + k + "}, HTTP: " + strconv.Itoa(status)
		return err
	}

	status = http.StatusOK
	k = "info"
	msg = "[INFO] SignUp completed: user signed up, email: {" + k + "}, HTTP: " + strconv.Itoa(status)
	return c.JSON(http.StatusOK, mc)

}

// revokes valid jwt token. Sends the token to Redis
func (j JWTHandler) SignOut(c echo.Context) error {
	var err error
	var status int = 200
	k, msg := "err", "[ERROR] "

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
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
		err = errors.New(msg)
		return err
	}

	// token already not valid
	if duration < 1 {
		msg += "SignOut error: duration lesser than 0, HTTP: " + strconv.Itoa(status)
		err = errors.New(msg)
		return err
	}

	// save token in Redis in order to blacklist it
	err = j.JwtC.JwtQE.SetToken(token, time.Duration(duration))
	if err != nil {
		status = http.StatusInternalServerError
		msg += "SignOut error: couldn't write to Redis, HTTP: " + strconv.Itoa(status)
		return err
	}

	fmt.Println("SIGNED  OUT")
	k = "info"
	msg = "[INFO] SignOut completed, HTTP: " + strconv.Itoa(status)
	return c.JSON(status, &model.GenericMessage{Message: msg})
}

// sign in a user
func (j JWTHandler) SignIn(c echo.Context) error {

	var authDetails Authentication
	var err error
	var status int
	k, msg := "err", "[ERROR]"

	defer func() {
		producer.ProduceMessage(k, msg)
		if err != nil {
			c.JSON(status, &model.GenericError{Message: msg})
		}
	}()

	if err = c.Bind(&authDetails); err != nil {
		status = http.StatusBadRequest
		k = authDetails.Email
		msg += "SignIn error: incorrect credentials, email: {" + k + "},  HTTP: " + strconv.Itoa(status)
		return err
	}

	// check if user exists
	var authUser model.User
	authUser, err = j.GetUser(authDetails.Email)

	k = authDetails.Email
	if err != nil {
		status = http.StatusInternalServerError
		msg += "SignIn error: user doesn't exist, email: {" + k + "},  HTTP: " + strconv.Itoa(status)
		return err
	}

	// check if password is correct
	check := j.JwtC.CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		status = http.StatusBadRequest
		msg += "SignIn error: incorrect password, email: {" + k + "},  HTTP: " + strconv.Itoa(status)
		err = errors.New("Incorrect password")
		return err
	}

	// generate token based on username and role
	var validToken string
	validToken, err = j.JwtC.GenerateJWT(authDetails.Email, authUser.Role)
	if err != nil {
		status = http.StatusInternalServerError
		msg += "SignIn error: couldn't generate token, email: {" + k + "},  HTTP: " + strconv.Itoa(status)
		return err
	}

	var token Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken
	status = http.StatusOK
	k = "info"
	msg = "[INFO] SignIn completed: user signed in, email: {" + k + "}, HTTP: " + strconv.Itoa(status)

	return c.JSON(status, token)
}
