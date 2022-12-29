package dto

import "final-project-backend/entity"

type ChangePasswordRequest struct {
	OldPassword string `binding:"required" json:"old_password"`
	NewPassword string `binding:"required" json:"new_password"`
}

type ChangePasswordResponse struct {
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
}




func (r *ChangePasswordResponse) BuildResponse(user entity.User) *ChangePasswordResponse {
	return &ChangePasswordResponse{
		Fullname: user.Fullname,
		Address:  user.Address,
	}
	}

