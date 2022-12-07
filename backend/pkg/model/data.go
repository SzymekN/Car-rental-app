package model

type DataModel interface{}

type Data struct{}

type Client struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PESEL       string `json:"pesel"`
	PhoneNumber string `json:"phone_number"`
	UserID      int    `json:"userId"`
	User        User   `json:"user"`
}

type Employee struct {
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
	RegistrationNumber string  `json:"registrationNumber,omitempty"`
	Brand              string  `json:"brand,omitempty"`
	Model              string  `json:"model,omitempty"`
	Type               string  `json:"type,omitempty"`
	Color              string  `json:"color,omitempty"`
	FuelConsumption    float32 `json:"fuelConsumption,omitempty"`
	DailyCost          int     `json:"dailyCost,omitempty"`
}

type GenericModel interface {
	User | Client | Employee | Vehicle
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
