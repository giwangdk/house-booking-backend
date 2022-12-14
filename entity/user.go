package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Photo    string `json:"photo"`
	Password string `json:"password"`
	Role     string `json:"role"`
	CityID   int    `json:"city_id"`
	City     City   `json:"city"`
}
