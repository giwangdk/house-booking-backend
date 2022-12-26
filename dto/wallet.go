package dto

import (
	"final-project-backend/entity"

	"github.com/shopspring/decimal"
)

type WalletDetail struct{
	ID int `json:"id"`
	Balance decimal.Decimal `json:"balance"`
	UserID int `json:"user_id"`
}

func (c *WalletDetail) BuildResponse(wallet entity.Wallet) *WalletDetail {
	return &WalletDetail{
		ID: int(wallet.ID),
		Balance: wallet.Balance,
		UserID: wallet.UserId,
	}
}