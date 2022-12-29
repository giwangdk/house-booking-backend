package dto

import (
	"final-project-backend/entity"
	"time"
)

type Reservation struct{
	ID int `json:"id"`
	CheckIn string `json:"check_in"`
	CheckOut string `json:"check_out"`
	TotalPrice int `json:"total_price"`
	Expired time.Time `json:"expired"`
	HouseID int `json:"house_id"`
	UserID int `json:"user_id"`
	StatusID int `json:"status_id"`
	BookingCode string `json:"booking_code"`
}

type ReservationList struct {
	Reservations []Reservation `json:"reservations"`
}

func (r *Reservation) BuildResponse(reservation entity.Reservation) *Reservation {
	return &Reservation{
		ID : int(reservation.ID),
		CheckIn: reservation.CheckIn,
		CheckOut: reservation.CheckOut,
		TotalPrice: reservation.TotalPrice,
		Expired: reservation.Expired,
		HouseID: reservation.HouseID,
		UserID: reservation.UserID,
		StatusID: reservation.StatusID,
		BookingCode: reservation.BookingCode,
	}
}

func (r *ReservationList) BuildResponse(reservations []*entity.Reservation) *ReservationList {
	var res []Reservation
	for _, reservation := range reservations {
		res = append(res, *(&Reservation{}).BuildResponse(*reservation))
	}
	return &ReservationList{
		Reservations: res,
	}
}