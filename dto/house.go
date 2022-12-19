package dto

import "final-project-backend/entity"

type House struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Description string `json:"description"`
	Location string `json:"location"`
	UserID int `json:"user_id"`
	City City `json:"city"`
}

func (c *House) BuildResponse(house entity.House) *House {
	city:= *(&City{}).BuildResponse(house.City)
	return &House{
		ID:   house.ID,
		Name: house.Name,
		Price: house.Price,
		Description: house.Description,
		Location: house.Location,
		UserID: house.UserID,
		City: city,	
	}
}