package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/repository"
)

type HouseUsecase interface {
	GetHouses() (*[]dto.House, error)
}

type HouseUsecaseImplementation struct {
	repository repository.HouseRepository
}

type HouseUsecaseImplementationConfig struct {
	Repository repository.HouseRepository
}

func NewHouseUseCase(c HouseUsecaseImplementationConfig) HouseUsecase {
	return &HouseUsecaseImplementation{
		repository: c.Repository,
	}
}

func (u *HouseUsecaseImplementation) GetHouses() (*[]dto.House, error) {

	var res []dto.House
	houses, err := u.repository.GetHouses()

	if err != nil {
		return nil, err
	}

	for _, house := range *houses {
		res = append(res, *(&dto.House{}).BuildResponse((house)))
	}

	return &res, nil
}
