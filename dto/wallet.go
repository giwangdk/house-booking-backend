package dto

import (
	"final-project-backend/entity"

	"github.com/shopspring/decimal"
)

type WalletDetail struct{
	ID uint `json:"id"`
	Balance decimal.Decimal `json:"balance"`
}

func (c *WalletDetail) BuildResponse(wallet entity.Wallet) *WalletDetail {
	return &WalletDetail{
		ID: wallet.ID,
		Balance: wallet.Balance,
	}
}