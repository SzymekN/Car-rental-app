package controller

import (
	"fmt"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// registers router for the server
func SetupRouter(svr *server.Server) {

	e := echo.New()
	svr.EchoServ = e

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	jwt_auth := e.Group("/api/v1")

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, `{"message":"Car sharing Welcome page!"}`)
	// })

	// create JWT handler and JWT validator config
	jwtH := auth.New(svr, svr.EchoServ, jwt_auth)
	jwtH.AddJWTMiddleware()

	// create all needed handlers
	authConf := auth.NewAuthConfig()
	uh := NewUserHandler(svr.GetMysqlDB(), svr.Logger, authConf, jwt_auth)

	// register all routes
	jwtH.RegisterRoutes()
	uh.RegisterRoutes()

	fmt.Println(jwtH)
	fmt.Println("setup", uh.SystemLogger)

}
