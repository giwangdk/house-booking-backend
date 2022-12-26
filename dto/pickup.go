package dto

import "final-project-backend/entity"

type Pickup struct {
	ReservationID  int `json:"reservation_id"`
	UserID         int `json:"user_id"`
	PickupStatusID int `json:"pickup_status_id"`
	PickupStatus   entity.PickupStatus
}
