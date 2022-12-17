package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/SzymekN/Car-rental-app/pkg/executor"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JWTHandler struct {
	JwtC  JWTControl
	group *echo.Group
}

type JWTHandlerInterface interface {
	RegisterRoutes()
	SignIn(c echo.Context) error
	SignUp(c echo.Context) error
	SignOut(c echo.Context) error
	revokeToken(token string, dur time.Duration)
}

func (j JWTHandler) RegisterRoutes() {
	j.JwtC.JwtQE.Svr.EchoServ.POST("/api/v1/users/signup", j.SignUp)
	j.JwtC.JwtQE.Svr.EchoServ.POST("/api/v1/users/signin", j.SignIn)
	j.group.GET("/users/signout", j.SignOut)

}

// Checks for the username in the db
func (j JWTHandler) getUserByEmail(email string) (model.User, producer.Log) {

	db := j.getMysqlDB()
	u := model.User{}
	log := producer.Log{}

	result := db.Debug().Where(&model.User{Email: email}).Find(&u)

	if err := result.Error; err != nil {
		code := http.StatusInternalServerError
		msg := fmt.Sprintf("[ERROR]: couldn't retrieve user, HTTP: %v", code)
		log.Populate("err", msg, code, err)
		return u, log
	}
	log = executor.CheckIfAffected(result)

	return u, log
}

// save signed in user to the db
func (j JWTHandler) SignUser(mc model.Client) producer.Log {
	db := j.getMysqlDB()
	log := producer.Log{}

	fmt.Println(db.Model(&model.Client{}).Association("User").Error)
	result := db.Model(&model.Client{}).Preload("User").Debug().Create(&mc)
	if err := result.Error; err != nil {
		code := http.StatusBadRequest
		msg := fmt.Sprintf("[ERROR]: couldn't create user, HTTP: %v", code)
		log.Populate("err", msg, code, err)
	}
	log = executor.CheckIfAffected(result)
	return log
}

func createSignUpResponse(mc model.Client, token string) SignInResponse {
	sir := SignInResponse{
		Email:       mc.User.Email,
		Name:        mc.Name,
		Surname:     mc.Surname,
		PhoneNumber: mc.PhoneNumber,
		Role:        "client",
		TokenString: token,
	}
	return sir
}

// Checks all passed credentials and saves user to the database
func (j JWTHandler) SignUp(c echo.Context) error {

	var validToken string
	logger := j.getLogger()
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("SignUp ")

	// user to save in the database
	var mc model.Client

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	// try saving data got in the request to the User datatype
	mc, logger.Log = executor.BindData(c, mc)
	if logger.Log.Err != nil {
		return logger.Log.Err
	}

	// check if user already exists
	_, logger.Log = j.getUserByEmail(mc.User.Email)
	if logger.Err != nil && logger.Err.Error() != "no rows affected" {
		return logger.Err
	}

	// hash password
	mc.User.Password, logger.Log = j.JwtC.GeneratehashPassword(mc.User.Password)
	if logger.Err != nil {
		return logger.Err
	}

	validToken, logger.Log = j.JwtC.GenerateJWT(mc.User.ID, mc.User.Email, mc.User.Role)

	if logger.Err != nil {
		return logger.Err
	}

	//insert user details to database
	mc.User.Role = "client"
	logger.Log = j.SignUser(mc)

	if logger.Err != nil {
		if logger.Err.Error() == "no rows affected" {
			code := http.StatusBadRequest
			msg := fmt.Sprintf("[ERROR]: duplicate entry - user exists, HTTP: %v", code)
			err := errors.New("duplicate entry")
			logger.Log.Populate("err", msg, code, err)
			return logger.Err
		} else {
			return logger.Err
		}
	}

	userResp := createSignUpResponse(mc, validToken)
	code := http.StatusOK
	k := "info"
	msg := "[INFO] SignUp completed: user signed up, email: {" + mc.User.Email + "}, HTTP: " + strconv.Itoa(code)
	logger.Populate(k, msg, code, nil)
	return c.JSON(code, userResp)

}

func (j JWTHandler) revokeToken(token string, dur time.Duration) {
	// j.JwtC.revokedTokens = append(j.JwtC.revokedTokens, token)
	go j.JwtC.JwtQE.SetToken(token, dur)
}

// revokes valid jwt token. Sends the token to Redis
func (j JWTHandler) SignOut(c echo.Context) error {
	logger := j.getLogger()
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("SignOut ")

	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
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
	} else if !ok {
		code := http.StatusBadRequest
		msg := "SignOut error: couldn't parse token, HTTP: " + strconv.Itoa(code)
		err := errors.New(msg)
		logger.Log.Populate("err", msg, code, err)
		return err
	}

	// token already not valid
	if duration < 1 {
		code := http.StatusBadRequest
		msg := "SignOut error: duration lesser than 0, HTTP: " + strconv.Itoa(code)
		err := errors.New(msg)
		logger.Log.Populate("err", msg, code, err)
		return err
	}

	// save token in Redis in order to blacklist it
	j.revokeToken(token, time.Duration(duration))

	// if err != nil {
	// 	status = http.StatusInternalServerError
	// 	msg += "SignOut error: couldn't write to Redis, HTTP: " + strconv.Itoa(status)
	// 	return err
	// }

	fmt.Println("SIGNED  OUT")
	code := http.StatusOK
	msg := "[INFO] SignOut completed, HTTP: " + strconv.Itoa(code)
	logger.Log.Populate("info", msg, code, nil)
	return c.JSON(code, &producer.GenericMessage{Message: msg})
}

// sign in a user
func (j JWTHandler) SignIn(c echo.Context) error {

	var authDetails Authentication
	logger := j.getLogger()
	logger.Log = producer.Log{}
	prefix := fmt.Sprintf("SignIn ")
	defer func() {
		logger.Log.Msg = fmt.Sprintf("%s %s", prefix, logger.Log.Msg)
		logger.ProduceWithJSON(c)
	}()

	if err := c.Bind(&authDetails); err != nil {
		code := http.StatusBadRequest
		msg := fmt.Sprintf("[ERROR]: couldn't bind data from request, HTTP: %v", code)
		logger.Populate("err", msg, code, err)
		return logger.Err
	}
	// check if user exists
	var authUser model.User
	authUser, logger.Log = j.getUserByEmail(authDetails.Email)

	if logger.Err != nil {
		return logger.Err
	}

	// check if password is correct
	logger.Log = CheckPasswordHash(authDetails.Password, authUser.Password)

	if logger.Err != nil {
		return logger.Err
	}

	// generate token based on username and role
	var validToken string

	validToken, logger.Log = j.JwtC.GenerateJWT(authUser.ID, authDetails.Email, authUser.Role)

	if logger.Err != nil {
		return logger.Err
	}

	db := j.getMysqlDB()

	var usrResponse SignInResponse
	switch authUser.Role {
	case "client":
		db.Debug().Table("client").Where("user_id = ?", authUser.ID).Find(&usrResponse)
	default:
		//TODO zmienić na employee
		db.Debug().Table("client").Where("user_id = ?", authUser.ID).Find(&usrResponse)
		// db.Debug().Table("employee").Where("user_id = ?", authUser.ID).Find(&usrResponse)
	}

	usrResponse.Email = authUser.Email
	usrResponse.Role = authUser.Role
	usrResponse.TokenString = validToken

	code := http.StatusOK
	msg := "[INFO] SignIn completed: user signed in, email: {" + usrResponse.Email + "}, HTTP: " + strconv.Itoa(code)
	logger.Log.Populate("info", msg, code, nil)
	return c.JSON(code, usrResponse)
}
