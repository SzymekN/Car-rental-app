package model

type Client struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	PESEL       string `json:"pesel,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	UserID      int    `json:"userId,omitempty"`
	User        User   `json:"user,omitempty"`
}

type Employee struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	PESEL       string `json:"pesel,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	UserID      int    `json:"userId,omitempty"`
	User        User   `json:"user,omitempty"`
}

type User struct {
	ID       int    `json:"id,omitempty"`
	Email    string `json:"email,omitempty" gorm:"email"`
	Password string `json:"password,omitempty" gorm:"password"`
	Role     string `json:"role,omitempty" gorm:"role"`
}

type Vehicle struct {
	ID                 int     `json:"id"`
	RegistrationNumber string  `json:"registrationNumber,omitempty"`
	Brand              string  `json:"brand,omitempty"`
	Model              string  `json:"model,omitempty"`
	Type               string  `json:"type,omitempty"`
	Color              string  `json:"color,omitempty"`
	FuelConsumption    float32 `json:"fuelConsumption,omitempty"`
	DailyCost          int     `json:"dailyCost,omitempty"`
}

type Salary struct {
	ID         int `json:"id"`
	Amount     int `json:"amount,omitempty"`
	EmployeeID int `json:"employee_id,omitempty"`
}

type Repairs struct {
	ID             int  `json:"id"`
	Cost           int  `json:"cost,omitempty"`
	Approved       bool `json:"approved,omitempty"`
	NotificationID int  `json:"notification_id,omitempty"`
	VehicleID      int  `json:"vehicle_id,omitempty"`
}

type Notification struct {
	ID          int    `json:"id"`
	Description string `json:"description,omitempty"`
	EmployeeID  int    `json:"employee_id,omitempty"`
	ClientID    int    `json:"client_id,omitempty"`
	VehicleID   int    `json:"vehicle_id,omitempty"`
}

type GenericModel interface {
	User | Client | Employee | Vehicle | Salary | Rental | Repairs | Notification
	GetId() int
}

func (d User) GetId() int {
	return d.ID
}

func (d Client) GetId() int {
	return d.ID
}

func (d Employee) GetId() int {
	return d.ID
}

func (d Vehicle) GetId() int {
	return d.ID
}

func (d Salary) GetId() int {
	return d.ID
}

func (d Notification) GetId() int {
	return d.ID
}
