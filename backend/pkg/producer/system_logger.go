package producer

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Log struct {
	Key  string
	Msg  string
	Code int
	Err  error
}

type GenericError struct {
	Message string `json:"message"`
}

type GenericMessage struct {
	Message string `json:"message"`
}

type SystemLogger struct {
	KafkaLogger
	Log
}
type LogProducer interface {
	Produce(c echo.Context)
}

func (l *Log) Populate(k, m string, c int, e error) {
	l.Key = k
	l.Msg = m
	l.Code = c
	l.Err = e
	fmt.Println(l)
}

func (sl SystemLogger) Produce(c echo.Context) {
	fmt.Println("SO: ", sl)
	if sl.Err != nil {
		sl.Msg += ", err:" + sl.Err.Error()
		c.JSON(sl.Code, &GenericMessage{Message: sl.Msg})
	}
	go sl.ProduceMessage(sl.Key, sl.Msg)
}
