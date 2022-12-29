package dto

type UpdateStatusPickupRequest struct {
	Status int `json:"status"`
}

type UpdateStatusPickupResponse struct {
	ID     int    `json:"id"`
	Status int `json:"status"`
}
