package producer

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type SystemLogger struct {
	KafkaLogger
	Log
}
type LogProducer interface {
	Produce(c echo.Context)
}

func (sl SystemLogger) ProduceLog() {
	fmt.Println("SO: ", sl)
	if sl.Err != nil {
		sl.Msg += ", err:" + sl.Err.Error()
	}
	go sl.ProduceMessage(sl.Key, sl.Msg)
}

func (sl SystemLogger) Produce(c echo.Context) {
	fmt.Println("SO: ", sl)
	if sl.Err != nil {
		sl.Msg += ", err:" + sl.Err.Error()
		c.JSON(sl.Code, &GenericMessage{Message: sl.Msg})
	}
	go sl.ProduceMessage(sl.Key, sl.Msg)
}
