package dto

import "final-project-backend/entity"

type CreateHouseDetailRequest struct {
	Bedrooms            int    `json:"bedrooms"`
	Beds                int    `json:"beds"`
	Baths               int    `json:"baths"`
	HouseFacilities     string `json:"house_facilities"`
	HouseRules          string `json:"house_rules"`
	HouseServices       string `json:"house_services"`
	BathroomsFacilities string `json:"bathrooms_facilities"`
	HouseID             int    `json:"house_id"`
}

type CreateHouseDetailResponse struct {
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
		Bedrooms:            houseDetail.Bedrooms,
		Beds:                houseDetail.Beds,
		Baths:               houseDetail.Baths,
		HouseFacilities:     houseDetail.HouseFacilities,
		HouseRules:          houseDetail.HouseRules,
		HouseServices:       houseDetail.HouseServices,
		BathroomsFacilities: houseDetail.BathroomsFacilities,
	}
}
