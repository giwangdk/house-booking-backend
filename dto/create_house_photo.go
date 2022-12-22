package dto

import "final-project-backend/entity"

type CreateHousePhotoRequest struct {
	HouseID int    `json:"house_id"`
	Photo   string `json:"photo" binding:"required"`
}

type CreateHousePhotoResponse struct {
	Photo string `json:"photo"`
}

func (c *CreateHousePhotoResponse) BuildResponse(house entity.HousePhoto) *CreateHousePhotoResponse {
	return &CreateHousePhotoResponse{
		Photo: house.Photo,
	}
}
