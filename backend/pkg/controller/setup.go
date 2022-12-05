package controller

import (
	"fmt"
	"net/http"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// registers router for the server
func SetupRouter(svr *server.Server) {
	// r := Router{e: echo.New()}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	// registerUserRoutes()

	jwtH := auth.New(svr)
	fmt.Println(jwtH)
	e.POST("/api/v1/users/signup", jwtH.SignUp)
	e.POST("/api/v1/users/signin", jwtH.SignIn)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"message":"Car sharing Welcome page!"}`)
	})

	// group of routes that will be validated with jwt
	jwt_auth := e.Group("/api/v1")
	config := jwtH.CreateJWTConfig()

	jwt_auth.Use(middleware.JWTWithConfig(config))

	jwt_auth.GET(" /users/signout", jwtH.SignOut)

	uc := UsersController{
		SystemOperator{
			DB:           svr.GetMysqlDB(),
			SystemLogger: svr.Logger,
		},
	}
	fmt.Println("setup", uc.SystemLogger)
	jwt_auth.GET("/users", uc.GetUserById)
	jwt_auth.GET("/users/all", uc.GetUsers)
	jwt_auth.POST("/users", uc.SaveUser, jwtH.JwtC.IsAdmin)
	jwt_auth.PUT("/users", uc.UpdateUser, jwtH.JwtC.IsAdmin)
	jwt_auth.DELETE("/users", uc.DeleteUser, jwtH.JwtC.IsAdmin)

	svr.EchoServ = e
}
