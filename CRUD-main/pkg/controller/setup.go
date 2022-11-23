package controller

import (
	"net/http"

	"github.com/SzymekN/CRUD/pkg/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mvrilo/go-redoc"
	echoredoc "github.com/mvrilo/go-redoc/echo"
)

func SetupRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"message":"Car sharing Welcome page!"}`)
	})

	e.POST("/api/v3/operators/signup", SignUp)
	e.GET("/api/v3/operators/signin", SignIn)

	jwt_auth := e.Group("")
	config := middleware.JWTConfig{
		SigningKey:     []byte(auth.Secretkey),
		ParseTokenFunc: auth.Validate,
	}

	jwt_auth.Use(middleware.JWTWithConfig(config))

	jwt_auth.GET("/api/v3/operators/signout", SignOut)

	jwt_auth.GET("/api/v1/users/:id", GetUserById)
	jwt_auth.GET("/api/v1/users", GetUsers)
	jwt_auth.POST("/api/v1/users/save", SaveUser, auth.IsAdmin)
	jwt_auth.PUT("/api/v1/users/:id", UpdateUser, auth.IsAdmin)
	jwt_auth.DELETE("/api/v1/users/:id", DeleteUser, auth.IsAdmin)

	// jwt_auth.GET("/api/v2/users/:id", GetUserByIdCassandraHandler)
	// jwt_auth.GET("/api/v2/users", GetUsersCassandraHandler)
	// jwt_auth.POST("/api/v2/users/save", SaveUserCassandraHandler, auth.IsAdmin)
	// jwt_auth.PUT("/api/v2/users/:id", UpdateUserCassandraHandler, auth.IsAdmin)
	// jwt_auth.DELETE("/api/v2/users/:id", DeleteUserCassandraHandler, auth.IsAdmin)

	// redoc documentation middleware
	doc := redoc.Redoc{
		Title:       "User API",
		Description: "API for interactions with database",
		SpecFile:    "docs/swagger.json",
		SpecPath:    "docs/swagger.json",
		DocsPath:    "/docs",
	}

	e.Use(echoredoc.New(doc))

	return e
}
