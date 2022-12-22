package dto

import "final-project-backend/entity"

type CreateHouseRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	UserID      int    `json:"user_id"`
	CityID      int    `json:"city_id" binding:"required"`
}

type CreateHouseResponse struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

func (c *CreateHouseResponse) BuildResponse(house entity.HouseProfile) *CreateHouseResponse {
	return &CreateHouseResponse{
		Name:        house.Name,
		Price:       house.Price,
		Description: house.Description,
		Location:    house.Location,
	}
}
