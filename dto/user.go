package dto

import "final-project-backend/entity"

type UserJWT struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}

type UserDetail struct {
	ID       uint    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Role     string `json:"role"`
	City    City `json:"city"`
}

type DetailUser struct {
	Fullname string      `json:"fullname"`
	Email    string      `json:"email"`
	Address  string      `json:"address"`
	Role     string      `json:"role"`
	CityID   int         `json:"city_id"`
	City    City `json:"city"`
}



func (c *UserDetail) BuildResponse(user entity.User) *UserDetail {
	city:= *(&City{}).BuildResponse(user.City)

	return &UserDetail{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email,
		Address:  user.Address,
		Role:     user.Role,
		City: city,		
	}
}
