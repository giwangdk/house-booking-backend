package dto

import "final-project-backend/entity"

type RegisterRequest struct {
	Fullname string `binding:"required"`
	Email    string `binding:"required,email"`
	Address  string `binding:"required"`
	Password string `binding:"required"`
	CityId   int    `binding:"required" json:"city_id"`
	
}

type RegisterResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	CityId     int    `json:"city_id"`
	Role     string `json:"role"`
}

func (r *RegisterRequest) BuildRequest(user RegisterRequest) *entity.User {
	return &entity.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Address:  user.Address,
	Password: user.Password,
	}
	}

func (r *RegisterResponse) BuildResponse(user entity.User) *RegisterResponse {
	return &RegisterResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
		Address:  user.Address,
		CityId:     user.CityID,
		Role:     user.Role,
	}
	}

