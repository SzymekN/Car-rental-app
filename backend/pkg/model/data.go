package model

import (
	"github.com/SzymekN/Car-rental-app/pkg/producer"
	"github.com/labstack/echo/v4"
)

type DataModel interface{}

type Data struct{}

// `Client` belongs to`User`, `UserID` is the foreign key
// type Client struct {
// 	ID          int    `json:"id" gorm:"->;primarykey"`
// 	Name        string `json:"name"`
// 	Surname     string `json:"surname"`
// 	PESEL       string `json:"pesel"`
// 	PhoneNumber string `json:"phoneNumber"`
// 	UserID      int    `json:"userId;omitempty"`
// 	User        `json:"user" gorm:"-;foreignKey:UserID;references:ID"`
// }

type Client struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PESEL       string `json:"pesel"`
	PhoneNumber string `json:"phone_number"`
	UserID      int    `json:"userId"`
	User        User   `json:"user"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
	Role     string `json:"role" gorm:"role"`
}

type Vehicle struct {
	ID                 int     `json:"id"`
	RegistrationNumber string  `json:"registrationNumber"`
	Brand              string  `json:"brand"`
	Model              string  `json:"model"`
	Type               string  `json:"type"`
	Color              string  `json:"color"`
	FuelConsumption    float32 `json:"fuelConsumption"`
	DailyCost          int     `json:"dailyCost"`
}

type Log struct {
	Key  string
	Msg  string
	Code int
	Err  error
}
type SystemLogger struct {
	producer.KafkaLogger
	Log
}
type LogProducer interface {
	Produce(c echo.Context)
}

func (sl SystemLogger) Produce(c echo.Context) {
	sl.ProduceMessage(sl.Key, sl.Msg)
	c.JSON(sl.Code, &GenericMessage{Message: sl.Msg})
}

type GenericModel interface {
	User | Client
	GetId() int
}

func (d User) GetId() int {
	return d.ID
}

func (d Client) GetId() int {
	return d.ID
}

// type RawJSON struct {
// 	Payload string `json:"payload"`
// }
// https://medium.com/cuddle-ai/building-microservice-using-golang-echo-framework-ff10ba06d508
// func (f *RawJSON) UnmarshalJSON(b []byte) error {
// 	type rawJ RawJSON
// 	newf := (*rawJ)(f)
// 	err := json.Unmarshal(b, newf)
// 	if err != nil {
// 		return err
// 	}

// 	var v DataModel
// 	err = json.Unmarshal([]byte(newf.Payload), &v)
// 	if err != nil {
// 		return err
// 	}
// 	var i interface{}
// 	switch v.(type) {
// 	case "user":
// 		i = &User{}
// 	default:
// 		return errors.New("unknown data type")
// 	}
// 	err = json.Unmarshal(raw, i)
// 	if err != nil {
// 		return err
// 	}
// 	f.Vehicles = append(f.Vehicles, i)

// 	return nil
// }

// swagger:model GenericError
type GenericError struct {
	// Response message
	// in: string
	// required: true
	Message string `json:"message"`
}

// swagger:model GenericMessage
type GenericMessage struct {
	// Response error with message
	// in: string
	// required: true
	Message string `json:"message"`
}
