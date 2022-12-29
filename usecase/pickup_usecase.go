package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
)

type PickupUsecase interface {
	CreatePickup(r dto.CreatePickupRequest) (*dto.CreatePickupResponse, error)
	UpdateStatusPickup(id int, statusID int) (*dto.CreatePickupResponse, error)
}

type PickupUsecaseImplementation struct {
	repository         repository.PickupRepository
}

type PickupUsecaseImplementationConfig struct {
	Repository         repository.PickupRepository
}

func NewPickupUseCase(c PickupUsecaseImplementationConfig) PickupUsecase {
	return &PickupUsecaseImplementation{
		repository:         c.Repository,
	}
}

func (u *PickupUsecaseImplementation) CreatePickup(r dto.CreatePickupRequest) (*dto.CreatePickupResponse, error) {

	pickup, err := u.repository.CreatePickup(entity.Pickup{
		ReservationID:  r.ReservationID,
		UserID:         r.UserID,
		PickupStatusID: 1,
	})
	if err != nil {
		return nil, err
	}

	res := (&dto.CreatePickupResponse{}).BuildResponse(*pickup)

	return res, nil

}

func (u *PickupUsecaseImplementation) UpdateStatusPickup(id int, statusID int) (*dto.CreatePickupResponse, error) {
	pickup, err := u.repository.UpdateStatus(id, statusID)
	if err != nil {
		return nil, err
	}
	res := (&dto.CreatePickupResponse{}).BuildResponse(*pickup)

	return res, nil
}
