package dto

import "final-project-backend/entity"

type UserJWT struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}

type DetailUser struct {
	Fullname string      `json:"fullname"`
	Email    string      `json:"email"`
	Address  string      `json:"address"`
	Role     string      `json:"role"`
	CityID   int         `json:"city_id"`
	City     entity.City `json:"city"`
	Password string      `json:"password"`
}

type EditUserRequest struct {
	Fullname    string `binding:"required" json:"fullname"`
	Address     string `binding:"required" json:"address"`
	OldPassword string `binding:"required" json:"old_password"`
	NewPassword string `binding:"required" json:"new_password"`
}

type EditUserResponse struct {
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
}
