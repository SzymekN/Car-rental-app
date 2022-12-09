package auth

import (
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"gorm.io/gorm"
)

type JWTWrappwer interface {
	getLogger() producer.SystemLogger
	ProduceMessage(k, val string)
	getMysqlDB() *gorm.DB
	getSigningKey() string
}

func (j JWTHandler) getLogger() producer.SystemLogger {
	return j.JwtC.JwtQE.Svr.Logger
}

// wrapper functions
func (j JWTHandler) ProduceMessage(k, val string) {
	j.JwtC.JwtQE.Svr.Logger.ProduceMessage(k, val)
}

func (j JWTHandler) getMysqlDB() *gorm.DB {
	return j.JwtC.JwtQE.Svr.GetMysqlDB()
}
func (j JWTHandler) getSigningKey() string {
	return j.JwtC.SecretKey
}
