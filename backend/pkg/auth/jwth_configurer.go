package auth

import (
	"context"

	"github.com/SzymekN/Car-rental-app/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTConfigurer interface {
	NewJWTHanlder(svr *server.Server, e *echo.Echo, g *echo.Group) *JWTHandler
	CreateJWTConfig() middleware.JWTConfig
	AddJWTMiddleware()
}

func NewJWTHanlder(svr *server.Server, e *echo.Echo, g *echo.Group) *JWTHandler {
	jwtH := &JWTHandler{
		JwtC: JWTControl{
			JwtQE: JWTQueryExecutor{
				Svr: svr,
				Ctx: context.Background(),
			},
			SecretKey: "",
		},
		echoServ: e,
		group:    g,
	}
	return jwtH
}

// Creates JWT configuration and adds middleware to group
func (j JWTHandler) AddJWTMiddleware() {
	config := j.CreateJWTConfig()
	j.group.Use(middleware.JWTWithConfig(config))
}

func (j JWTHandler) CreateJWTConfig() middleware.JWTConfig {
	conf := middleware.JWTConfig{
		SigningKey:     []byte(j.getSigningKey()),
		ParseTokenFunc: j.JwtC.Validate,
	}
	return conf
}
