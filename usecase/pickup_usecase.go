package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/httperror"
	"final-project-backend/repository"
)

type PickupUsecase interface {
	GetPickups(page int, limit int, sortBy string, sort string, searchBy string, filterByStatus int) (*dto.PickupLists, error)
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

func (u *PickupUsecaseImplementation) GetPickups(page int, limit int, sortBy string, sort string, searchBy string, filterByStatus int) (*dto.PickupLists, error) {
	pickups,total,  err := u.repository.GetPickups(page, limit, sortBy, sort, searchBy, filterByStatus)
	if err != nil {
		return nil, err
	}

	res := (&dto.PickupLists{}).BuildResponse(pickups, page, limit, total)

	return res, nil
}

func (u *PickupUsecaseImplementation) CreatePickup(r dto.CreatePickupRequest) (*dto.CreatePickupResponse, error) {

	pickup, err := u.repository.CreatePickup(entity.Pickup{
		ReservationID:  r.ReservationID,
		UserID:         r.UserID,
		PickupStatusID: 1,
	})
	if err != nil {
		return nil, httperror.BadRequestError("Failed to create pickup", "FAILED_CREATE_PICKUP")
	}

	res := (&dto.CreatePickupResponse{}).BuildResponse(*pickup)

	return res, nil

}

func (u *PickupUsecaseImplementation) UpdateStatusPickup(id int, statusID int) (*dto.CreatePickupResponse, error) {
	pickup, err := u.repository.UpdateStatus(id, statusID)
	if err != nil {
		return nil, httperror.BadRequestError("Failed to update status", "ERROR_UPDATE_STATUS")
	}
	res := (&dto.CreatePickupResponse{}).BuildResponse(*pickup)

	return res, nil
}
