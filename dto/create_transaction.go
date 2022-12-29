package dto

import (
	"final-project-backend/entity"
	"mime/multipart"
)

type CreateTransactionRequest struct {
	BookingCode string `json:"booking_code" binding:"required"`
	IsGuest bool `json:"is_guest"`
	TransferSlip multipart.File  `json:"transfer_slip" form:"file" binding:"required"`
}

type CreateTransactionResponse struct {
	ReservationID int `json:"reservation_id"`
	HouseID int `json:"house_id"`
	UserID int `json:"user_id"`
}


type CreateTransactionGuestResponse struct {
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

