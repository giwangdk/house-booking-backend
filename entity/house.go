package entity

import "gorm.io/gorm"

type House struct {
	gorm.Model
	Name string `json:"name"`
	Price int `json:"price"`
	Description string `json:"description"`
	Location string `json:"location"`
	UserID int `json:"user_id"`
	CityID int `json:"city_id"`
	City    City `json:"city"`
}

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

type HousePhoto struct {
	gorm.Model
	Photo string `json:"photo"`
	HouseID int `json:"house_id"`
}