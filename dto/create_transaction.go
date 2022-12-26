package dto

import "final-project-backend/entity"

type CreateTransactionRequest struct {
	ReservationID int `json:"reservation_id"`
	BookingCode string `json:"booking_code"`
}

type CreateTransactionResponse struct {
	ReservationID int `json:"reservation_id"`
	HouseID int `json:"house_id"`
	UserID int `json:"user_id"`
}

func (r *CreateTransactionResponse) BuildResponse(transaction entity.Transaction) *CreateTransactionResponse {
	return &CreateTransactionResponse{
		ReservationID: transaction.ReservationID,
		HouseID: transaction.HouseID,
		UserID: transaction.UserID,
	}
}