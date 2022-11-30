package model

type DataModel interface{}

type Data struct{}

// `Client` belongs to`User`, `UserID` is the foreign key
// type Client struct {
// 	ID          int    `json:"id" gorm:"->;primarykey"`
// 	Name        string `json:"name"`
// 	Surname     string `json:"surname"`
// 	PESEL       string `json:"pesel"`
// 	PhoneNumber string `json:"phoneNumber"`
// 	UserID      int    `json:"userId;omitempty"`
// 	User        `json:"user" gorm:"-;foreignKey:UserID;references:ID"`
// }

type Client struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PESEL       string `json:"pesel"`
	PhoneNumber string `json:"phoneNumber"`
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
	RegistrationNumber string  `json:"registrationNumber"`
	Brand              string  `json:"brand"`
	Model              string  `json:"model"`
	Type               string  `json:"type"`
	Color              string  `json:"color"`
	FuelConsumption    float32 `json:"fuelConsumption"`
	DailyCost          int     `json:"dailyCost"`
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

// swagger:model GenericError
type GenericError struct {
	// Response message
	// in: string
	// required: true
	Message string `json:"message"`
}

// swagger:model GenericMessage
type GenericMessage struct {
	// Response error with message
	// in: string
	// required: true
	Message string `json:"message"`
}
