package entity

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Game struct{
	gorm.Model
	Chance decimal.Decimal `json:"chance"`
	UserId int `json:"user_id"`
}