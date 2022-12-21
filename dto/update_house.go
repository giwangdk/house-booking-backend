package dto

import "final-project-backend/entity"

type UpdateHouseRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Location    string `json:"location"`
	UserID      int    `json:"user_id"`
	CityID      int    `json:"city_id"`
}

type UpdateHouseResponse struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

func (r *UpdateHouseResponse) BuildResponse(entity entity.House) *UpdateHouseResponse {
	return &UpdateHouseResponse{
		Name:        entity.Name,
		Price:       entity.Price,
		Description: entity.Description,
		Location:    entity.Location,
	}
}
