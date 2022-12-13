package model

import "google.golang.org/genproto/googleapis/type/datetime"

type Rental struct {
	ID            int               `json:"id"`
	StartDate     datetime.DateTime `json:"start_date,omitempty"`
	EndDate       datetime.DateTime `json:"end_date,omitempty"`
	PickupAddress string            `json:"pickup_address,omitempty"`
	EmployeeID    int               `json:"employee_id,omitempty"`
	ClientID      int               `json:"client_id,omitempty"`
	VehicleID     int               `json:"vehicle_id,omitempty"`
}

func (d Rental) GetId() int {
	return d.ID
}
