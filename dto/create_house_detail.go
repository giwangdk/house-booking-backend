package dto

import (
	"final-project-backend/entity"

	"github.com/shopspring/decimal"
)

type CreateHouseDetailRequest struct {
	MaxGuest            decimal.Decimal    `json:"max_guest" binding:"required"`
	Bedrooms            decimal.Decimal    `json:"bedrooms" binding:"required"`
	Beds                decimal.Decimal    `json:"beds" binding:"required"`
	Baths               decimal.Decimal    `json:"baths" binding:"required"`
	HouseFacilities     string `json:"house_facilities" binding:"required"`
	HouseRules          string `json:"house_rules"`
	HouseServices       string `json:"house_services"`
	BathroomsFacilities string `json:"bathrooms_facilities"`
	HouseID             int   `json:"house_id"`
}

type CreateHouseDetailResponse struct {
	MaxGuest            decimal.Decimal    `json:"max_guest"`
	Bedrooms            decimal.Decimal    `json:"bedrooms"`
	Beds                decimal.Decimal    `json:"beds"`
	Baths               decimal.Decimal    `json:"baths"`
	HouseFacilities     string `json:"house_facilities"`
	HouseRules          string `json:"house_rules"`
	HouseServices       string `json:"house_services"`
	BathroomsFacilities string `json:"bathrooms_facilities"`
}

func (c *CreateHouseDetailResponse) BuildResponse(houseDetail entity.HouseDetail) *CreateHouseDetailResponse {
	return &CreateHouseDetailResponse{
		MaxGuest:            houseDetail.MaxGuest,
		Bedrooms:            houseDetail.Bedrooms,
		Beds:                houseDetail.Beds,
		Baths:               houseDetail.Baths,
		HouseFacilities:     houseDetail.HouseFacilities,
		HouseRules:          houseDetail.HouseRules,
		HouseServices:       houseDetail.HouseServices,
		BathroomsFacilities: houseDetail.BathroomsFacilities,
	}
}
