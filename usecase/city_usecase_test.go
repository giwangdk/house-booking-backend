package usecase_test

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"final-project-backend/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCities(t *testing.T) {
	t.Run("Should return list of city", func(t *testing.T) {

		cityRepo := new(mocks.CityRepository)

		uc := usecase.NewCityUseCase(usecase.CityUsecaseImplementationConfig{
			Repository: cityRepo,
		})

		var cities []entity.City

		cityRepo.On("GetCities").Return(&cities, nil)
		res, err := uc.GetCities()

		assert.Nil(t, err)
		assert.Equal(t, cities, *res)
	})

	t.Run("Should return error when get cities", func(t *testing.T) {
		
		cityRepo := new(mocks.CityRepository)

		uc := usecase.NewCityUseCase(usecase.CityUsecaseImplementationConfig{
			Repository: cityRepo,
		})

		cityRepo.On("GetCities").Return(nil, errors.New("error"))
		res, err := uc.GetCities()

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})



	}

