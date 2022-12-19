package dto

import "final-project-backend/entity"

type EditUserRequest struct {
	Fullname    string `binding:"required" json:"fullname"`
	Address     string `binding:"required" json:"address"`
	OldPassword string `binding:"required" json:"old_password"`
	NewPassword string `binding:"required" json:"new_password"`
}

type EditUserResponse struct {
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
}

type UpdateRoleRequest  struct {
	Email   string `binding:"required" json:"email"`
}


func (r *EditUserResponse) BuildResponse(user entity.User) *EditUserResponse {
	return &EditUserResponse{
		Fullname: user.Fullname,
		Address:  user.Address,
	}
	}

