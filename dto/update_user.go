package dto

import "final-project-backend/entity"

type UpdateUserRequest struct {
	Fullname    string `binding:"required" json:"fullname"`
	Address     string `binding:"required" json:"address"`
	OldPassword string `binding:"required" json:"old_password"`
	NewPassword string `binding:"required" json:"new_password"`
}

type UpdateUserResponse struct {
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
}

type UpdateRoleRequest  struct {
	Email   string `binding:"required" json:"email"`
}


func (r *UpdateUserResponse) BuildResponse(user entity.User) *UpdateUserResponse {
	return &UpdateUserResponse{
		Fullname: user.Fullname,
		Address:  user.Address,
	}
	}

