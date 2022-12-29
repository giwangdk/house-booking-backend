package dto

import (
	"final-project-backend/entity"

	"github.com/shopspring/decimal"
)

type UpdateHouseDetailRequest struct {
	MaxGuest            decimal.Decimal    `json:"max_guest"`
	Bedrooms            decimal.Decimal    `json:"bedrooms"`
	Beds                decimal.Decimal    `json:"beds"`
	Baths               decimal.Decimal    `json:"baths"`
	HouseFacilities     string `json:"house_facilities"`
	HouseRules          string `json:"house_rules"`
	HouseServices       string `json:"house_services"`
	BathroomsFacilities string `json:"bathrooms_facilities"`
	HouseID             int  `json:"house_id"`
}

type UpdateHouseDetailResponse struct {
	MaxGuest            decimal.Decimal    `json:"max_guest"`
	Bedrooms            decimal.Decimal    `json:"bedrooms"`
	Beds                decimal.Decimal    `json:"beds"`
	Baths               decimal.Decimal    `json:"baths"`
	HouseFacilities     string `json:"house_facilities"`
	HouseRules          string `json:"house_rules"`
	HouseServices       string `json:"house_services"`
	BathroomsFacilities string `json:"bathrooms_facilities"`
}

func (r *UpdateHouseDetailResponse) BuildResponse(entity entity.HouseDetail) *UpdateHouseDetailResponse {
	return &UpdateHouseDetailResponse{
		MaxGuest:            entity.MaxGuest,
		Bedrooms:            entity.Bedrooms,
		Beds:                entity.Beds,
		Baths:               entity.Baths,
		HouseFacilities:     entity.HouseFacilities,
		HouseRules:          entity.HouseRules,
		HouseServices:       entity.HouseServices,
		BathroomsFacilities: entity.BathroomsFacilities,
	}
}
