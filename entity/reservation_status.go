package entity

import "gorm.io/gorm"

type ReservationStatus struct {
	gorm.Model
	Status string `json:"status"`
}