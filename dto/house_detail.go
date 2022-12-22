package dto

import "final-project-backend/entity"

type HouseDetail struct {
	MaxGuest            int    `json:"max_guest"`
	Bedrooms            int    `json:"bedrooms"`
	Beds                int    `json:"beds"`
	Baths               int    `json:"baths"`
	HouseFacilities     string `json:"house_facilities"`
	HouseRules          string `json:"house_rules"`
	HouseServices       string `json:"house_services"`
	BathroomsFacilities string `json:"bathrooms_facilities"`
}

func (h *HouseDetail) BuildResponse(entity entity.HouseDetail) *HouseDetail {
	return &HouseDetail{
		Bedrooms:        entity.Bedrooms,
		Beds:            entity.Beds,
		Baths:           entity.Baths,
		HouseFacilities: entity.HouseFacilities,
		HouseRules:      entity.HouseRules,
		HouseServices:   entity.HouseServices,
	}
}
