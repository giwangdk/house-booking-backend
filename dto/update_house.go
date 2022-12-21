package dto

type UpdateHouseRequest struct {
	Fullname    string `binding:"required" json:"fullname"`
	Address     string `binding:"required" json:"address"`
	OldPassword string `binding:"required" json:"old_password"`
	NewPassword string `binding:"required" json:"new_password"`
}

type UpdateHouseResponse struct {
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
}




