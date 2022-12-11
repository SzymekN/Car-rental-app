package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// key used for singing jwt tokens

type JWTControl struct {
	JwtQE     JWTQueryExecutor
	SecretKey string
}

type JWTControllerInterface interface {
	GeneratehashPassword(password string) (string, producer.Log)
	CheckPasswordHash(password, hash string) producer.Log
	Validate(auth string, c echo.Context) (interface{}, error)
	GenerateJWT(email, role string) (string, producer.Log)
}

// func (j JWTControl) checkToken(val string) (bool, error) {

// 	for _, revokedToken := range j.revokedTokens {
// 		if revokedToken == val {
// 			return true, nil
// 		}
// 	}

//		// revoked, _ := j.JwtQE.GetToken(val)
//		return false, nil
//	}
func (j JWTControl) produceMessage(k, val string) {
	j.JwtQE.Svr.Logger.ProduceMessage(k, val)
}

func (j JWTControl) GeneratehashPassword(password string) (string, producer.Log) {
	log := producer.Log{}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		code := http.StatusInternalServerError
		msg := fmt.Sprintf("[ERROR]: password hashing failure, HTTP: %v", code)
		log.Populate("err", msg, code, err)
	}
	return string(bytes), log
}

func (j JWTControl) CheckPasswordHash(password, hash string) producer.Log {
	log := producer.Log{}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		code := http.StatusUnauthorized
		msg := fmt.Sprintf("[ERROR]: wrong password, HTTP: %v", code)
		log.Populate("err", msg, code, err)
	}
	return log
}

// check if token is correct
func (j JWTControl) Validate(auth string, c echo.Context) (interface{}, error) {

	// method of obtaining validating key from the database
	remoteKeyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}

		// query Redis for the key
		j.SecretKey = j.getKey()
		return []byte(j.SecretKey), nil
	}

	// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
	token, err := jwt.Parse(auth, remoteKeyFunc)
	// check if this token is already revoked
	// error possible
	tokenRevoked, _ := j.JwtQE.GetToken(token.Raw)

	if tokenRevoked {
		go j.produceMessage("JWT validation", token.Raw+" REVOKED")
		return nil, errors.New("Token Revoked")
	}

	// check if errors occured during token generation
	if err != nil {
		go j.produceMessage("JWT validation", "JWT validation failed: "+err.Error())
		return nil, err
	}

	// check if generated token is valid
	if !token.Valid {
		go j.produceMessage("JWT validation", "JWT validation failed: invalid token")
		return nil, errors.New("invalid token")
	}
	go j.produceMessage("JWT validation", "JWT validation succesfull")
	return token, nil
}

// set signing key for application instance
func (j JWTControl) getKey() string {
	var err error

	// check if key is already set, if not query Redis for it
	if j.SecretKey == "" {
		j.SecretKey, err = j.JwtQE.getSigningKey()
	}
	// if key doesn't exist in Redis, generate it
	if err != nil {
		j.SecretKey, err = j.JwtQE.setSigningKey()
	}

	if err != nil {
		panic("No jwt Key found")
	}
	return j.SecretKey
}

// generates valid token based on username, role and expire date
func (j JWTControl) GenerateJWT(email, role string) (string, producer.Log) {
	var mySigningKey = []byte(j.getKey())
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expireTime := time.Minute * 60
	log := producer.Log{}

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(expireTime).Unix()

	// sign created token
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		code := http.StatusInternalServerError
		msg := fmt.Sprintf("[ERROR]: JWT generation failure, HTTP: %v", code)
		log.Populate("err", msg, code, err)
	}
	return tokenString, log
}
