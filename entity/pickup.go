package entity

import "gorm.io/gorm"

type Pickup struct {
	gorm.Model
	ReservationID  int          `json:"reservation_id"`
	UserID         int          `json:"user_id"`
	PickupStatusID int          `json:"pickup_status_id"`
	PickupStatus   PickupStatus `json:"pickup_status"`
}
