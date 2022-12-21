package dto

import "final-project-backend/entity"

type HousePhotoDetail struct {
	ID    int    `json:"id"`
	Photo string `json:"photo"`
}

type HousePhotoLists struct {
	HousePhotos []HousePhotoDetail `json:"house_photos"`
}

func (h *HousePhotoDetail) BuildResponse(entity entity.HousePhoto) *HousePhotoDetail {
	return &HousePhotoDetail{
		ID:    int(entity.ID),
		Photo: entity.Photo,
	}
}

func (h *HousePhotoLists) BuildResponse(entities []entity.HousePhoto) *[]HousePhotoDetail {
	var res []HousePhotoDetail
	for _, entity := range entities {
		res = append(res, *(&HousePhotoDetail{}).BuildResponse(entity))
	}
	return &res
}
