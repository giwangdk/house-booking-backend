package dto

import "final-project-backend/entity"

type CreateHouseRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Location    string `json:"location"`
	UserID      int    `json:"user_id"`
	CityID      int    `json:"city_id"`
}

type CreateHouseResponse struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

func (c *CreateHouseResponse) BuildResponse(house entity.House) *CreateHouseResponse {
	return &CreateHouseResponse{
		Name:        house.Name,
		Price:       house.Price,
		Description: house.Description,
		Location:    house.Location,
	}
}
