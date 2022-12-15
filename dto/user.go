package dto

import "final-project-backend/entity"

type UserJWT struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type DetailUser struct {
	Fullname string      `json:"fullname"`
	Email    string      `json:"email"`
	Address  string      `json:"address"`
	Role     string      `json:"role"`
	CityID   int         `json:"city_id"`
	City     entity.City `json:"city"`
}
