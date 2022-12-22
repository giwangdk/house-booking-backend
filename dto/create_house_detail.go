package dto

import "final-project-backend/entity"

type CreateHouseDetailRequest struct {
	MaxGuest            int    `json:"max_guest" binding:"required"`
	Bedrooms            int    `json:"bedrooms" binding:"required"`
	Beds                int    `json:"beds" binding:"required"`
	Baths               int    `json:"baths" binding:"required"`
	HouseFacilities     string `json:"house_facilities" binding:"required"`
	HouseRules          string `json:"house_rules"`
	HouseServices       string `json:"house_services"`
	BathroomsFacilities string `json:"bathrooms_facilities"`
	HouseID             int    `json:"house_id"`
}

type CreateHouseDetailResponse struct {
	MaxGuest            int    `json:"max_guest"`
	Bedrooms            int    `json:"bedrooms"`
	Beds                int    `json:"beds"`
	Baths               int    `json:"baths"`
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
