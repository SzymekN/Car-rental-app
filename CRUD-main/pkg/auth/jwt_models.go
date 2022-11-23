package auth

type Operator struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

type Authentication struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Token struct {
	Role        string `json:"role" form:"role"`
	Username    string `json:"username" form:"username"`
	TokenString string `json:"token" form:"token"`
}
