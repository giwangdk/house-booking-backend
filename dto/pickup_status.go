package dto

import "final-project-backend/entity"

type PickupStatus struct {
	ID   int    `json:"id"`
	Status string `json:"status"`
}

func (p *PickupStatus) BuildResponse(pickup entity.PickupStatus) *PickupStatus {
	return &PickupStatus{
		ID:     int(pickup.ID),
		Status: pickup.Status,
	}
}