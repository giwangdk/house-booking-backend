package dto

import (
	"final-project-backend/entity"

	"github.com/shopspring/decimal"
)

type HouseDetail struct {
	ID 			   uint  `json:"id"`
	MaxGuest            decimal.Decimal   `json:"max_guest"`
	Bedrooms            decimal.Decimal   `json:"bedrooms"`
	Beds                decimal.Decimal   `json:"beds"`
	Baths               decimal.Decimal   `json:"baths"`
	HouseFacilities     string `json:"house_facilities"`
	HouseRules          string `json:"house_rules"`
	HouseServices       string `json:"house_services"`
	BathroomsFacilities string `json:"bathrooms_facilities"`
}

func (h *HouseDetail) BuildResponse(entity entity.HouseDetail) *HouseDetail {
	return &HouseDetail{
		ID:              entity.ID,
		Bedrooms:        entity.Bedrooms,
		Beds:            entity.Beds,
		Baths:           entity.Baths,
		HouseFacilities: entity.HouseFacilities,
		HouseRules:      entity.HouseRules,
		HouseServices:   entity.HouseServices,
	}
}
