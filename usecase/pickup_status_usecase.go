package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/repository"
)

type PickupStatusUsecase interface {
	GetPickupStatus() (*dto.PickupStatusLists, error)
}

type PickupStatusUsecaseImplementation struct {
	repository         repository.PickupStatusRepository
}

type PickupStatusUsecaseImplementationConfig struct {
	Repository         repository.PickupStatusRepository
}

func NewPickupStatusUseCase(c PickupStatusUsecaseImplementationConfig) PickupStatusUsecase {
	return &PickupStatusUsecaseImplementation{
		repository:         c.Repository,
	}
}

func (u *PickupStatusUsecaseImplementation) GetPickupStatus() (*dto.PickupStatusLists, error) {
	pickupStatus,  err := u.repository.GetPickupStatus()
	if err != nil {
		return nil, err
	}

	res := (&dto.PickupStatusLists{}).BuildResponse(pickupStatus)

	return res, nil
}
