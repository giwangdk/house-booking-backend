package dto

import "final-project-backend/entity"

type UpdateHouseDetailRequest struct {
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

type UpdateHouseDetailResponse struct {
	MaxGuest            int    `json:"max_guest"`
	Bedrooms            int    `json:"bedrooms"`
	Beds                int    `json:"beds"`
	Baths               int    `json:"baths"`
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
