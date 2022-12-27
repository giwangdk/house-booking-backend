package dto

import "final-project-backend/entity"

type House struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	Price       int                `json:"price"`
	Description string             `json:"description"`
	Location    string             `json:"location"`
	User        UserDetail         `json:"user"`
	City        City               `json:"city"`
	Photos      []HousePhotoDetail `json:"photos"`
	HouseDetail `json:"detail"`
}

type HouseProfile struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Location    string `json:"location"`
	UserID      int    `json:"user_id"`
	CityID      int    `json:"city_id"`
}

type HouseLists struct {
	Houses []House `json:"houses"`
	Page   int     `json:"page"`
	Limit  int     `json:"limit"`
	Total  int     `json:"total"`
}

func (c *House) BuildResponse(house entity.House) *House {
	user := *(&UserDetail{}).BuildResponse(house.User)
	city := *(&City{}).BuildResponse(house.City)
	photos := *(&HousePhotoLists{}).BuildResponse(house.Photos)
	return &House{
		ID:          house.ID,
		Name:        house.Name,
		Price:       house.Price,
		Description: house.Description,
		Location:    house.Location,
		User:        user,
		City:        city,
		Photos:      photos,
		HouseDetail: *(&HouseDetail{}).BuildResponse(house.HouseDetail),
	}
}

func (c *HouseLists) BuildResponse(houses []entity.House, page int, limit int, total int) *HouseLists {
	var res []House
	for _, house := range houses {
		res = append(res, *(&House{}).BuildResponse(house))
	}
	return &HouseLists{
		Houses: res,
		Page:   page,
		Limit:  limit,
		Total:  total,
	}
}

func (c *HouseProfile) BuildResponse(house entity.HouseProfile) *HouseProfile {
	return &HouseProfile{
		Name:        house.Name,
		Price:       house.Price,
		Description: house.Description,
		Location:    house.Location,
		UserID:      house.UserID,
		CityID:      house.CityID,
	}
}
