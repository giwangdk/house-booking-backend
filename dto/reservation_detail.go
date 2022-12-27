package dto

import (
	"final-project-backend/entity"
	"time"
)

type ReservationDetail struct {
	ID int `json:"id"`
	CheckIn string `json:"check_in"`
	CheckOut string `json:"check_out"`
	TotalPrice int `json:"total_price"`
	Expired time.Time `json:"expired"`
	HouseID int `json:"house_id"`
	UserID int `json:"user_id"`
	StatusID int `json:"status_id"`
	BookingCode string `json:"booking_code"`
	House HouseProfile `json:"house"`
	User UserDetail `json:"user"`
}

func (r *ReservationDetail) BuildResponse(reservation entity.Reservation) *ReservationDetail {

	user := *(&UserDetail{}).BuildResponse(reservation.User)
	house:= *(&HouseProfile{}).BuildResponse(reservation.House)

	return &ReservationDetail{
		ID: int(reservation.ID),
		CheckIn: reservation.CheckIn,
		CheckOut: reservation.CheckOut,
		TotalPrice: reservation.TotalPrice,
		Expired: reservation.Expired,
		HouseID: reservation.HouseID,
		UserID: reservation.UserID,
		StatusID: reservation.StatusID,
		BookingCode: reservation.BookingCode,
		House: house,
		User: user,
	}
}