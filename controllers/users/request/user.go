package request

type login_request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type register_request struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
