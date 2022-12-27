package entity

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct{
	gorm.Model
	CheckIn string `json:"check_in"`
	CheckOut string `json:"check_out"`
	TotalPrice int `json:"total_price"`
	Expired time.Time `json:"expired"`
	HouseID int `json:"house_id"`
	UserID int `json:"user_id"`
	StatusID int `json:"status_id"`
	BookingCode string `json:"booking_code"`
	House HouseProfile  `json:"house" gorm:"foreignKey:HouseID"`
	User User  `json:"user" gorm:"foreignKey:UserID"`
}