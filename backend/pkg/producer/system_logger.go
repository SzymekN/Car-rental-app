package producer

import (
	"github.com/labstack/echo/v4"
)

type SystemLogger struct {
	KafkaLogger
	Log
}
type LogProducer interface {
	ProduceWithJSON(c echo.Context)
	Produce(c echo.Context)
}

func (sl SystemLogger) ProduceLog() {
	if sl.Err != nil {
		sl.Msg += ", err:" + sl.Err.Error()
	}
	go LogToDB(sl.Key, sl.Msg)
	go sl.ProduceMessage(sl.Key, sl.Msg)
}

func (sl SystemLogger) ProduceWithJSON(c echo.Context) {
	if sl.Err != nil {
		sl.Msg += ", err:" + sl.Err.Error()
		c.JSON(sl.Code, &GenericMessage{Message: sl.Msg})
	}
	go LogToDB(sl.Key, sl.Msg)
	go sl.ProduceMessage(sl.Key, sl.Msg)
}
func (sl SystemLogger) Produce() {
	go LogToDB(sl.Key, sl.Msg)
	go sl.ProduceMessage(sl.Key, sl.Msg)
}
