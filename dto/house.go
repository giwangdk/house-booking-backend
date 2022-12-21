package dto

import "final-project-backend/entity"

type House struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Description string `json:"description"`
	Location string `json:"location"`
	User UserDetail `json:"user"`
	City City `json:"city"`
}

type HouseLists struct {
	Houses []House `json:"houses"`
	Page int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}
func (c *House) BuildResponse(house entity.House) *House {
	user := *(&UserDetail{}).BuildResponse(house.User)
	city:= *(&City{}).BuildResponse(house.City)
	return &House{
		ID:   house.ID,
		Name: house.Name,
		Price: house.Price,
		Description: house.Description,
		Location: house.Location,
		User: user,
		City: city,	
	}
}

func (c *HouseLists) BuildResponse(houses []entity.House, page int, limit int, total int) *HouseLists {
	var res []House
	for _, house := range houses {
		res = append(res, *(&House{}).BuildResponse(house))
	}
	return &HouseLists{
		Houses: res,
		Page: page,
		Limit: limit,
		Total: total,
	}
}
