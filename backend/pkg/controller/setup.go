package controller

import (
	"net/http"

	"github.com/SzymekN/Car-rental-app/pkg/auth"
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type MainController struct {
}

type Controller interface {
	GetDB() *gorm.DB
}

type Router struct {
	e *echo.Echo
}

type Registrator interface {
	GetEcho() *echo.Echo
	registerUserRoutes()
}

func (mc MainController) GetDB() *gorm.DB {
	return storage.MysqlConn.GetDBInstance()
}

func (r Router) registerUserRoutes() {
	e := r.e
	e.POST("/api/v1/users/signup", SignUp)
	e.POST("/api/v1/users/signin", SignIn)

}

func test(c echo.Context) error {
	return GenericPost(c, model.User{})
}

// registers router for the server
func SetupRouter() *echo.Echo {
	// r := Router{e: echo.New()}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	// registerUserRoutes()
	e.POST("/api/v1/users/signup", SignUp)
	e.POST("/api/v1/users/signin", SignIn)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"message":"Car sharing Welcome page!"}`)
	})

	// group of routes that will be validated with jwt
	jwt_auth := e.Group("")
	config := middleware.JWTConfig{
		SigningKey:     []byte(auth.Secretkey),
		ParseTokenFunc: auth.Validate,
	}

	jwt_auth.Use(middleware.JWTWithConfig(config))

	jwt_auth.GET("/api/v1/users/signout", SignOut)

	uc := UsersController{}
	jwt_auth.GET("/api/v1/user", uc.GetUserById)
	jwt_auth.GET("/api/v1/users", uc.GetUsers)
	jwt_auth.POST("/api/v1/users/save", uc.SaveUser, auth.IsAdmin)
	jwt_auth.PUT("/api/v1/users/:id", uc.UpdateUser, auth.IsAdmin)
	jwt_auth.DELETE("/api/v1/users/:id", uc.DeleteUser, auth.IsAdmin)

	e.POST("/test", test)

	// redoc documentation middleware
	// doc := redoc.Redoc{
	// 	Title:       "User API",
	// 	Description: "API for interactions with database",
	// 	SpecFile:    "docs/swagger.json",
	// 	SpecPath:    "docs/swagger.json",
	// 	DocsPath:    "/docs",
	// }

	// e.Use(echoredoc.New(doc))

	return e
}
