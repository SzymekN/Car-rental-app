package controller

import (
	"fmt"
	"net/http"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// type MainController struct {
// 	e  *echo.Echo
// 	uc UsersController
// }

// type Controller interface {
// 	RegisterRoutes()
// 	GetDB() *gorm.DB
// }

// func (mc MainController) GetDB() *gorm.DB {
// 	return storage.MysqlConn.GetDBInstance()
// }

// func (mc MainController) RegisterRoutes() {

// }

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
	jwt_auth := e.Group("")
	config := middleware.JWTConfig{
		SigningKey:     []byte(jwtH.GetSigningKey()),
		ParseTokenFunc: jwtH.JwtC.Validate,
	}

	jwt_auth.Use(middleware.JWTWithConfig(config))

	jwt_auth.GET("/api/v1/users/signout", jwtH.SignOut)

	uc := UsersController{db: svr.GetMysqlDB()}
	jwt_auth.GET("/api/v1/users", uc.GetUserById)
	jwt_auth.GET("/api/v1/users/all", uc.GetUsers)
	jwt_auth.POST("/api/v1/users", uc.SaveUser, jwtH.JwtC.IsAdmin)
	jwt_auth.PUT("/api/v1/users", uc.UpdateUser, jwtH.JwtC.IsAdmin)
	jwt_auth.DELETE("/api/v1/users", uc.DeleteUser, jwtH.JwtC.IsAdmin)

	svr.E = e
}
