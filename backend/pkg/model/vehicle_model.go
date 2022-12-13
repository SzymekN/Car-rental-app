package model

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

func (d Vehicle) GetId() int {
	return d.ID
}
