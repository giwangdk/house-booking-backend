package dto

import "final-project-backend/entity"

type Pickup struct {
	ID 		   int `json:"id"`
	ReservationID  int `json:"reservation_id"`
	UserID         int `json:"user_id"`
	PickupStatusID int `json:"pickup_status_id"`
	PickupStatus   entity.PickupStatus `json:"pickup_status"`
}

type PickupLists struct {
	Pickups []Pickup `json:"pickups"`
	Page   int     `json:"page"`
	Limit  int     `json:"limit"`
	Total  int     `json:"total"`
}



func (p *Pickup) BuildResponse(pickup entity.Pickup) *Pickup {
	return &Pickup{
		ID:             int(pickup.ID),
		ReservationID:  pickup.ReservationID,
		UserID:         pickup.UserID,
		PickupStatusID: pickup.PickupStatusID,
		PickupStatus:   pickup.PickupStatus,
	}
}

func (p *PickupLists) BuildResponse(pickups []entity.Pickup, page int, limit int, total int) *PickupLists {
	var pickupList []Pickup
	for _, pickup := range pickups {
		pickupList = append(pickupList, Pickup{
			ID:             int(pickup.ID),
			ReservationID:  pickup.ReservationID,
			UserID:         pickup.UserID,
			PickupStatusID: pickup.PickupStatusID,
			PickupStatus:   pickup.PickupStatus,
		})
	}
	return &PickupLists{
		Pickups: pickupList,
		Page:    page,
		Limit:   limit,
		Total:   total,
	}
}