package dto

import (
	"final-project-backend/entity"
	"mime/multipart"
)

type CreateHousePhotoRequest struct {
	HouseID int    `json:"house_id"`
	Photo   multipart.File   `json:"photo" form:"file" binding:"required"`
}

type CreateHousePhotoResponse struct {
	Photo string `json:"photo"`
}

func (c *CreateHousePhotoResponse) BuildResponse(house entity.HousePhoto) *CreateHousePhotoResponse {
	return &CreateHousePhotoResponse{
		Photo: house.Photo,
	}
}
