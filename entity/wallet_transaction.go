package entity

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type WalletTransaction struct {
	gorm.Model
	Sender      int64           `json:"sender"`
	Amount      decimal.Decimal `json:"amount"`
	Recipient   int64           `json:"recipient"`
	Description string          `json:"description"`
}
