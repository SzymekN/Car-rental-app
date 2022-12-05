package server

import (
	"github.com/SzymekN/Car-rental-app/pkg/storage"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	E         *echo.Echo
	MysqlConn storage.MysqlConnect
	RedisConn storage.RedisConnect
}

// var svr Server

func (svr Server) GetMysqlDB() *gorm.DB {
	return svr.MysqlConn.DB
}
func (svr Server) GetRedisDB() *redis.Client {
	return svr.RedisConn.RDB
}
