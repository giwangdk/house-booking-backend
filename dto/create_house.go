package dto

import (
	"final-project-backend/entity"

	"github.com/shopspring/decimal"
)

type CreateHouseRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       decimal.Decimal    `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	UserID      int    `json:"user_id"`
	CityID      int    `json:"city_id" binding:"required"`
}

type CreateHouseResponse struct {
	ID 		int    `json:"id"`
	Name        string `json:"name"`
	Price       decimal.Decimal    `json:"price"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

func (c *CreateHouseResponse) BuildResponse(house entity.HouseProfile) *CreateHouseResponse {
	return &CreateHouseResponse{
		ID: 		int(house.ID),
		Name:        house.Name,
		Price:       house.Price,
		Description: house.Description,
		Location:    house.Location,
	}
}
