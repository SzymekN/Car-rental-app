package auth

// user data type and fields
type Operator struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

// login credentials
type Authentication struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// parts of a jwt token, combined and hashed to create token
type Token struct {
	Role        string `json:"role" form:"role"`
	Username    string `json:"username" form:"username"`
	TokenString string `json:"token" form:"token"`
}
