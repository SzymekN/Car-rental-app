package controller

import (
	"fmt"

	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/labstack/echo/v4"
)

type BasicHandler interface {
	RegisterRoutes()
	Save(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetById(c echo.Context) error
	GetAll(c echo.Context) error
}

func HandleRequestResult[T any](c echo.Context, data T, log producer.Log) error {
	if log.Err != nil {
		log.Msg += ", err:" + log.Err.Error()
		fmt.Println(log)
		return c.JSON(log.Code, &producer.GenericMessage{Message: log.Msg})
	} else {
		return c.JSON(log.Code, data)
	}
}
