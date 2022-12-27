package model

type Employee struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	PESEL   string `json:"pesel,omitempty"`
	Salary  int    `json:"salary,omitempty"`
	UserID  int    `json:"userId,omitempty"`
	User    User   `json:"user,omitempty"`
}

func (d Employee) GetId() int {
	return d.ID
}
