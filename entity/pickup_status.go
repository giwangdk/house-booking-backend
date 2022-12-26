package entity

import "gorm.io/gorm"

type PickupStatus struct {
	gorm.Model
	Status string `json:"status"`
}