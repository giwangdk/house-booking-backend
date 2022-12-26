package dto

import "final-project-backend/entity"

type CreatePickupRequest struct {
	ReservationID int `json:"reservation_id" binding:"required"`
	UserID        int `json:"user_id" binding:"required"`
}

type CreatePickupResponse struct {
	ReservationID int                 `json:"reservation_id"`
	UserID        int                 `json:"user_id"`
	PickupStatus  entity.PickupStatus `json:"pickup_status"`
}

func (r *CreatePickupResponse) BuildResponse(pickup entity.Pickup) *CreatePickupResponse {
	return &CreatePickupResponse{
		ReservationID: pickup.ReservationID,
		UserID:        pickup.UserID,
		PickupStatus:  pickup.PickupStatus,
	}
}
