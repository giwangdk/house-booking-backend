package dto

import "final-project-backend/entity"

type CreateReservationRequest struct {
	CheckIn string `json:"check_in"`
	CheckOut string `json:"check_out"`
	TotalPrice int `json:"total_price"`
}

type CreateReservationResponse struct {
	CheckIn string `json:"check_in"`
	CheckOut string `json:"check_out"`
	TotalPrice int `json:"total_price"`
	HouseID int `json:"house_id"`
	StatusID int `json:"status_id"`
}

func (r *CreateReservationResponse) BuildResponse(reservation entity.Reservation) *CreateReservationResponse {
	return &CreateReservationResponse{
		CheckIn: reservation.CheckIn,
		CheckOut: reservation.CheckOut,
		TotalPrice: reservation.TotalPrice,
		HouseID: reservation.HouseID,
		StatusID: reservation.StatusID,
	}
}