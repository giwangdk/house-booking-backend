package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ReservationID int `json:"reservation_id"`
	UserID int `json:"user_id"`
	HouseID int `json:"house_id"`
	TransferSlip string `json:"transfer_slip"`
}