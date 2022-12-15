package model

import (
	"time"
)

type Rental struct {
	ID            int       `json:"id"`
	StartDate     time.Time `json:"start_date,omitempty"`
	EndDate       time.Time `json:"end_date,omitempty"`
	PickupAddress string    `json:"pickup_address,omitempty"`
	DriverID      *int      `json:"employee_id,omitempty" default:"nil"`
	ClientID      int       `json:"client_id,omitempty"`
	VehicleID     int       `json:"vehicle_id,omitempty"`
}

func (d Rental) GetId() int {
	return d.ID
}
