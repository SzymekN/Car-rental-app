package model

type Notification struct {
	ID          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	EmployeeID  *int   `json:"employee_id,omitempty" default:"nil"`
	ClientID    *int   `json:"client_id,omitempty" default:"nil"`
	VehicleID   *int   `json:"vehicle_id,omitempty" default:"nil"`
}

func (d Notification) GetId() int {
	return d.ID
}
