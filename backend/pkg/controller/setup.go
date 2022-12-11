package controller

import (
	"net/http"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/SzymekN/Car-rental-app/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// type MainController struct {
// 	jwtH     *auth.JWTHandler
// 	// uH       *UserHandler
// 	// vH       *VehicleHandler
// 	handlers []BasicController
// }

// func (mc *MainController) RegisterAllRoutes() {
// 	for _, handler := range mc.handlers {
// 		handler.RegisterRoutes()
// 	}
// }

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
	// create all needed handlers
	authConf := auth.NewAuthConfig()
	// mc.uH = NewUserHandler(systemOperator, authConf, jwt_auth)
	// mc.vH = NewVehicleHandler(systemOperator, authConf, jwt_auth)
	mc.handlers = append(mc.handlers, NewUserHandler(systemOperator, authConf, jwt_auth))
	mc.handlers = append(mc.handlers, NewVehicleHandler(systemOperator, authConf, jwt_auth))

	// register all routes
	mc.RegisterAllRoutes()
	// mc.jwtH.RegisterRoutes()
	// mc.uH.RegisterRoutes()
	// mc.vH.RegisterRoutes()

	// fmt.Println(mc.jwtH)
	// fmt.Println("setup", mc.uH.sysOperator.SystemLogger)

}
