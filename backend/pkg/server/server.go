package server

import (
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"github.com/SzymekN/Car-rental-app/pkg/storage"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	EchoServ  *echo.Echo
	MysqlConn storage.MysqlConnect
	RedisConn storage.RedisConnect
	Logger    model.SystemLogger
}

// var svr Server

func (svr Server) GetMysqlDB() *gorm.DB {
	return svr.MysqlConn.GetDb()
}

func (svr Server) GetRedisDB() *redis.Client {
	return svr.RedisConn.GetRDB()
}

// func (svr Server) GetSystemLogger() *redis.Client {
// 	return svr.RedisConn.GetRDB()
// }
