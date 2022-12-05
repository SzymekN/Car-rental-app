package controller

import (
	"github.com/SzymekN/Car-rental-app/pkg/model"
	"gorm.io/gorm"
)

type SystemOperator struct {
	DB *gorm.DB
	model.SystemLogger
}

func (sl SystemOperator) GetDB() *gorm.DB {
	return sl.DB
}

// func (sl SystemOperator) GetSystemLogger() model.SystemLogger {
// 	return sl.SystemLogger
// }
