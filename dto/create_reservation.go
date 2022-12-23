package dto

import (
	"final-project-backend/entity"
	"time"
)

type CreateReservationRequest struct {
	CheckIn string `json:"check_in" binding:"required"`
	CheckOut string `json:"check_out" binding:"required"`
	TotalPrice int `json:"total_price" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Email string `json:"email" binding:"required"`
	CityID int `json:"city_id" binding:"required"`
	HouseID int `json:"house_id" binding:"required"`

}

type CreateReservationResponse struct {
	CheckIn string `json:"check_in"`
	CheckOut string `json:"check_out"`
	TotalPrice int `json:"total_price"`
	HouseID int `json:"house_id"`
	StatusID int `json:"status_id"`
	UserID int `json:"user_id"`
	Expire time.Time `json:"expire"`
}

func (r *CreateReservationResponse) BuildResponse(reservation entity.Reservation) *CreateReservationResponse {
	return &CreateReservationResponse{
		CheckIn: reservation.CheckIn,
		CheckOut: reservation.CheckOut,
		TotalPrice: reservation.TotalPrice,
		HouseID: reservation.HouseID,
		StatusID: reservation.StatusID,
		UserID: reservation.UserID,
		Expire: reservation.Expired,
	}
}