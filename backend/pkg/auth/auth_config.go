package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type pathPrivileges map[string][]string

type AuthConfig struct {
	//array of maps where key is path and value is a string with authorized roles
	Privileges pathPrivileges
}

type AuthConfigInterface interface {
	Contains(path, role string) bool
	NewAuthConfig() AuthConfig
	IsAuthorized(next echo.HandlerFunc) echo.HandlerFunc
}

func (ac AuthConfig) Contains(path, role string) bool {

	s, ok := ac.Privileges[path]
	if !ok {
		return false
	}

	for _, v := range s {
		if v == role {
			return true
		}
	}
	return false
}

func NewAuthConfig() AuthConfig {
	conf := AuthConfig{
		Privileges: pathPrivileges{
			"/api/v1/users":        {"owner"},
			"/api/v1/users/all":    {"owner"},
			"/api/v1/clients/self": {"client"},
		},
	}
	return conf
}

// checks for role embedded in the token to get information about privileges
func (ac AuthConfig) IsAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		path := c.Path()
		fmt.Println(c.Path())
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		if role == "admin" {
			fmt.Println("AUTH: ADMIN")
			return next(c)
		} else if ac.Contains(path, role) {
			fmt.Println("AUTH: ", role)
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}
