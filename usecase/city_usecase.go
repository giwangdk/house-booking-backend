package usecase

import (
	"final-project-backend/entity"
	"final-project-backend/repository"
)

type CityUsecase interface {
	GetCities() (*[]entity.City, error)
}

type CityUsecaseImplementation struct {
	repository repository.CityRepository
}

type CityUsecaseImplementationConfig struct {
	Repository repository.CityRepository
}

func NewCityUseCase(c CityUsecaseImplementationConfig) CityUsecase {
	return &CityUsecaseImplementation{
		repository: c.Repository,
	}
}

func (u *CityUsecaseImplementation) GetCities() (*[]entity.City, error) {
	cities, err := u.repository.GetCities()

	if err != nil {
		return nil, err
	}

	return cities, nil
}
