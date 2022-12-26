package model

import (
	"fmt"
	"strings"
	"time"
)

type DateRange struct {
	StartDate CustomTime `json:"start_date,omitempty"`
	EndDate   CustomTime `json:"end_date,omitempty"`
}

type CustomTime struct {
	time.Time
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	date := t.Time.Format("2006-01-02")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}
