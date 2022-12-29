package dto

type UpdateStatusPickupRequest struct {
	PickupStatusID int `json:"pickup_status_id"`
}

type UpdateStatusPickupResponse struct {
	ID     int    `json:"id"`
	PickupStatusID int `json:"pickup_status_id"`
}
