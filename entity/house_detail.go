package entity

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type HouseDetail struct {
	gorm.Model
	MaxGuest           decimal.Decimal   `json:"max_guest"`
	Bedrooms           decimal.Decimal   `json:"bedrooms"`
	Beds               decimal.Decimal   `json:"beds"`
	Baths              decimal.Decimal   `json:"baths"`
	HouseFacilities     string `json:"house_facilities"`
	HouseRules          string `json:"house_rules"`
	HouseServices       string `json:"house_services"`
	BathroomsFacilities string `json:"bathrooms_facilities"`
	HouseID            int   `json:"house_id"`
}
