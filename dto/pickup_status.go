package dto

import "final-project-backend/entity"

type PickupStatus struct {
	ID   int    `json:"id"`
	Status string `json:"status"`
}

type PickupStatusLists struct {
	Data []PickupStatus `json:"data"`
}

func (p *PickupStatus) BuildResponse(pickup entity.PickupStatus) *PickupStatus {
	return &PickupStatus{
		ID:     int(pickup.ID),
		Status: pickup.Status,
	}
}

func (p *PickupStatusLists) BuildResponse(pickupStatus []entity.PickupStatus) *PickupStatusLists {
	var pickupStatusList []PickupStatus
	for _, pickup := range pickupStatus {
		pickupStatusList = append(pickupStatusList, *(&PickupStatus{}).BuildResponse(pickup))
	}

	return &PickupStatusLists{
		Data: pickupStatusList,
	}
}
