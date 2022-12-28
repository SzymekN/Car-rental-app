package model

type User struct {
	ID       int    `json:"id,omitempty"`
	Email    string `json:"email,omitempty" gorm:"email"`
	Password string `json:"password,omitempty" gorm:"password"`
	Role     string `json:"role,omitempty" gorm:"role"`
}

func (d User) GetId() int {
	return d.ID
}

type UIDWrapper struct {
	UID int `json:"id"`
}
