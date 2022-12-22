package dto

import "final-project-backend/entity"

type UpdateHouseRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	CityID      int    `json:"city_id" binding:"required"`
}

type UpdateHouseResponse struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

func (r *UpdateHouseResponse) BuildResponse(entity entity.HouseProfile) *UpdateHouseResponse {
	return &UpdateHouseResponse{
		Name:        entity.Name,
		Price:       entity.Price,
		Description: entity.Description,
		Location:    entity.Location,
	}
}
