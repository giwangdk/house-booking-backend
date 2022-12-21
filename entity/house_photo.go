package entity

import "gorm.io/gorm"

type HousePhoto struct {
	gorm.Model
	HouseID int    `json:"house_id"`
	Photo string `json:"photo"`
}