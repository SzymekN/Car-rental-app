package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/SzymekN/CRUD/pkg/producer"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var Secretkey string = ""

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Validate(auth string, c echo.Context) (interface{}, error) {

	remoteKeyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}

		Secretkey = getKey()
		return []byte(Secretkey), nil
	}

	// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
	token, err := jwt.Parse(auth, remoteKeyFunc)
	tokenRevoked, _ := GetToken(token.Raw)

	if tokenRevoked {
		producer.ProduceMessage("JWT validation", token.Raw+" REVOKED")
		return nil, errors.New("Token Revoked")
	}

	// zwrócony token i nil == poprawny token
	if err != nil {
		producer.ProduceMessage("JWT validation", "JWT validation failed: "+err.Error())
		return nil, err
	}
	if !token.Valid {
		producer.ProduceMessage("JWT validation", "JWT validation failed: invalid token")
		return nil, errors.New("invalid token")
	}

	producer.ProduceMessage("JWT validation", "JWT validation succesfull")
	return token, nil
}

func getKey() string {
	var err error
	if Secretkey == "" {
		Secretkey, err = getSigningKey()
	}
	if err != nil {
		Secretkey, err = setSigningKey()
	}

	if err != nil {
		panic("No jwt Key found")
	}
	return Secretkey
}

func GenerateJWT(username, role string) (string, error) {
	var mySigningKey = []byte(getKey())
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expireTime := time.Minute * 2

	claims["authorized"] = true
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(expireTime).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"]
		if role != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
