package model

type Notification struct {
	ID          int    `json:"id"`
	Description string `json:"description,omitempty"`
	EmployeeID  int    `json:"employee_id,omitempty"`
	ClientID    int    `json:"client_id,omitempty"`
	VehicleID   int    `json:"vehicle_id,omitempty"`
}

func (d Notification) GetId() int {
	return d.ID
}
