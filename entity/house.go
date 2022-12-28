package entity

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type HouseProfile struct {
	gorm.Model
	Name        string `json:"name"`
	Price       decimal.Decimal    `json:"price"`
	Description string `json:"description"`
	Location    string `json:"location"`
	UserID      int    `json:"user_id"`
	CityID      int    `json:"city_id"`
	City        City   `json:"city"`
	User        User   `json:"user"`
}

type House struct {
	gorm.Model
	Name        string       `json:"name"`
	Price       decimal.Decimal          `json:"price"`
	Description string       `json:"description"`
	Location    string       `json:"location"`
	UserID      int          `json:"user_id"`
	CityID      int          `json:"city_id"`
	City        City         `json:"city"`
	User        User         `json:"user"`
	Photos      []HousePhoto `json:"photos"`
	HouseDetail `json:"detail"`
}

func (HouseProfile) TableName() string {
	return "houses"
}
