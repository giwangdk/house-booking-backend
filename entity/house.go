package entity

import "gorm.io/gorm"

type House struct {
	gorm.Model
	Name        string       `json:"name"`
	Price       int          `json:"price"`
	Description string       `json:"description"`
	Location    string       `json:"location"`
	UserID      int          `json:"user_id"`
	CityID      int          `json:"city_id"`
	City        City         `json:"city"`
	User        User         `json:"user"`
	Photos      []HousePhoto `json:"photos"`
	HouseDetail `json:"detail"`
}
