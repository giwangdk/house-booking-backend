package dto

type RegisterRequest struct {
	Fullname string `binding:"required"`
	Email    string `binding:"required,email"`
	Address  string `binding:"required"`
	Password string `binding:"required"`
	CityId   int    `binding:"required"`
}

type RegisterResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	City     int    `json:"city_id"`
	Role     string `json:"role"`
}
