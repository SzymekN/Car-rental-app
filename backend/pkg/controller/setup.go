package controller

import (
	"net/http"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/SzymekN/Car-rental-app/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var E *echo.Echo

// registers router for the server
func SetupRouter(svr *server.Server) {

	e := echo.New()
	mc := MainController{}
	svr.EchoServ = e

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	jwt_auth := e.Group("/api/v1")

	systemOperator := producer.NewSystemOperator(svr.GetMysqlDB(), svr.Logger)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"message":"Car sharing Welcome page!"}`)
	})

	// create JWT handler and JWT validator config
	mc.jwtH = auth.NewJWTHandler(svr, jwt_auth)
	mc.jwtH.AddJWTMiddleware()
	authConf := auth.NewAuthConfig()

	// create all needed handlers

	E = e
	mc.handlers = append(mc.handlers, NewClientHandler(systemOperator, authConf, jwt_auth))
	mc.handlers = append(mc.handlers, NewEmployeeHandler(systemOperator, authConf, jwt_auth))
	mc.handlers = append(mc.handlers, NewNotificationHandler(systemOperator, authConf, jwt_auth))
	mc.handlers = append(mc.handlers, NewRepairHandler(systemOperator, authConf, jwt_auth))
	mc.handlers = append(mc.handlers, NewSalaryHandler(systemOperator, authConf, jwt_auth))
	mc.handlers = append(mc.handlers, NewUserHandler(systemOperator, authConf, jwt_auth))
	mc.handlers = append(mc.handlers, NewVehicleHandler(systemOperator, authConf, jwt_auth))
	mc.handlers = append(mc.handlers, NewRentalHandler(systemOperator, authConf, jwt_auth))
	mc.handlers = append(mc.handlers, NewLogHandler(systemOperator, authConf, jwt_auth))

	mc.RegisterAllRoutes()

	// data, err := json.MarshalIndent(e.Routes(), "", "  ")
	// if err != nil {
	// 	return
	// }
	// fmt.Println(string(data))

}
