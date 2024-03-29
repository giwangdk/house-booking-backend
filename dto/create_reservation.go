package dto

import (
	"final-project-backend/entity"
	"time"

	"github.com/shopspring/decimal"
)

type CreateReservationRequest struct {
	CheckIn         string `json:"check_in" binding:"required"`
	CheckOut        string `json:"check_out" binding:"required"`
	TotalPrice      decimal.Decimal    `json:"total_price" binding:"required"`
	Fullname        string `json:"fullname" binding:"required"`
	Email           string `json:"email" binding:"required"`
	CityID          int    `json:"city_id" binding:"required"`
	HouseID         int    `json:"house_id" binding:"required"`
	IsRequestPickup bool   `json:"is_request_pickup" binding:"required"`
}

type CreateReservationResponse struct {
	CheckIn     string    `json:"check_in"`
	CheckOut    string    `json:"check_out"`
	TotalPrice  decimal.Decimal       `json:"total_price"`
	HouseID     int       `json:"house_id"`
	StatusID    int       `json:"status_id"`
	UserID      int       `json:"user_id"`
	Expire      time.Time `json:"expire"`
	BookingCode string    `json:"booking_code"`
}

func (r *CreateReservationResponse) BuildResponse(reservation entity.Reservation) *CreateReservationResponse {
	return &CreateReservationResponse{
		CheckIn:     reservation.CheckIn,
		CheckOut:    reservation.CheckOut,
		TotalPrice:  reservation.TotalPrice,
		HouseID:     reservation.HouseID,
		StatusID:    reservation.StatusID,
		UserID:      reservation.UserID,
		Expire:      reservation.Expired,
		BookingCode: reservation.BookingCode,
	}
}
