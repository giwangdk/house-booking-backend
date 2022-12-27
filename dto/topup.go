package dto

import "github.com/shopspring/decimal"

type TopUpRequest struct {
	Sender    int
	Recipient int
	Amount    decimal.Decimal `json:"amount" binding:"required"`
}

type TopUpResponse struct {
	Sender      int
	Amount      decimal.Decimal
	Recipient   int
	Description string
}
