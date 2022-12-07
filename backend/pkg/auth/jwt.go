package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// key used for singing jwt tokens

type JWTControl struct {
	JwtQE     JWTQueryExecutor
	SecretKey string
}

func (j JWTControl) ProduceMessage(k, val string) {
	j.JwtQE.Svr.Logger.ProduceMessage(k, val)
}

func (j JWTControl) GeneratehashPassword(password string) (string, error) {
	fmt.Printf("HASHED PASSWORD:%s\n", password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (j JWTControl) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
	tokenRevoked, _ := j.JwtQE.GetToken(token.Raw)

	if tokenRevoked {
		j.ProduceMessage("JWT validation", token.Raw+" REVOKED")
		return nil, errors.New("Token Revoked")
	}

	// check if errors occured during token generation
	if err != nil {
		j.ProduceMessage("JWT validation", "JWT validation failed: "+err.Error())
		return nil, err
	}

	// check if generated token is valid
	if !token.Valid {
		j.ProduceMessage("JWT validation", "JWT validation failed: invalid token")
		return nil, errors.New("invalid token")
	}

	j.ProduceMessage("JWT validation", "JWT validation succesfull")
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
func (j JWTControl) GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(j.getKey())
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	expireTime := time.Minute * 15

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(expireTime).Unix()

	// sign created token
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
