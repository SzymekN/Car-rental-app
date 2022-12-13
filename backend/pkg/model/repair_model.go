package model

type Repair struct {
	ID             int  `json:"id"`
	Cost           int  `json:"cost,omitempty"`
	Approved       bool `json:"approved,omitempty"`
	NotificationID int  `json:"notification_id,omitempty"`
	VehicleID      int  `json:"vehicle_id,omitempty"`
}

func (d Repair) GetId() int {
	return d.ID
}
