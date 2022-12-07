package auth

// user data type and fields
// type User struct {
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// 	Role     string `json:"role" form:"role"`
// }

// login credentials
type Authentication struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// parts of a jwt token, combined and hashed to create token
type SignInResponse struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role" gorm:"role"`
	TokenString string `json:"token"`
}
