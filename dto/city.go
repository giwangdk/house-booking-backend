package dto

import "final-project-backend/entity"

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *City) BuildResponse(city entity.City) *City {
	return &City{
		ID:   city.Id,
		Name: city.Name,
	}
}