package entity

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Wallet struct{
	gorm.Model
	Balance decimal.Decimal `json:"balance"`
	UserId int `json:"user_id"`
}