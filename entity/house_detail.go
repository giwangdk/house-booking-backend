package entity

import "gorm.io/gorm"



type HouseDetail struct {
	gorm.Model
	Bedrooms int `json:"bedrooms"`
	Beds int `json:"beds"`
	Baths int `json:"baths"`
	HouseFacilities  string `json:"house_facilities"`
	HouseRules string `json:"house_rules"`
	HouseServices string `json:"house_services"`
	BathroomsFacilities string `json:"bathrooms_facilities"`
	HouseID int `json:"house_id"`
}