package producer

import (
	"gorm.io/gorm"
)

type SystemOperator struct {
	DB *gorm.DB
	SystemLogger
}

func (sl SystemOperator) GetDB() *gorm.DB {
	return sl.DB
}

func NewSystemOperator(db *gorm.DB, l SystemLogger) SystemOperator {
	so := SystemOperator{
		DB:           db,
		SystemLogger: l,
	}
	return so
}

// func (sl SystemOperator) GetSystemLogger() producer.SystemLogger {
// 	return sl.SystemLogger
// }
