package model

type Salary struct {
	ID         int `json:"id,omitempty"`
	Amount     int `json:"amount,omitempty"`
	EmployeeID int `json:"employee_id,omitempty"`
}

func (d Salary) GetId() int {
	return d.ID
}
