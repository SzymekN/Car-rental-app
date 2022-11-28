package auth

// user data type and fields
type User struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

// login credentials
type Authentication struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// parts of a jwt token, combined and hashed to create token
type Token struct {
	Role        string `json:"role" form:"role"`
	Email       string `json:"email" form:"email"`
	TokenString string `json:"token" form:"token"`
}
