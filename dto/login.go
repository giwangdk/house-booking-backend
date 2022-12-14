package dto

type LoginRequest struct {
	Email    string `binding:"required,email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
