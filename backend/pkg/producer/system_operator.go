package producer

import (
	"fmt"
	"strings"
	"time"

	"github.com/SzymekN/Car-rental-app/pkg/model"
	"gorm.io/gorm"
)

type SystemOperator struct {
	DB *gorm.DB
	SystemLogger
}

func (sl SystemOperator) GetDB() *gorm.DB {
	return sl.DB
}

var global_db *gorm.DB

func LogToDB(k, v string) {
	k = strings.TrimSpace(k)
	v = strings.TrimSpace(v)
	l := model.Log{
		Timestamp: time.Now(),
		Key:       k,
		Value:     v,
	}
	if err := global_db.Create(&l).Error; err != nil {
		fmt.Println("Error while inserting log")
	}
}

func NewSystemOperator(db *gorm.DB, l SystemLogger) SystemOperator {
	so := SystemOperator{
		DB:           db,
		SystemLogger: l,
	}
	global_db = so.DB
	return so
}

// func (sl SystemOperator) GetSystemLogger() producer.SystemLogger {
// 	return sl.SystemLogger
// }
