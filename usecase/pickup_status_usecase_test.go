package usecase_test

import (
	"errors"
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"final-project-backend/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPickupStatus(t *testing.T) {
	t.Run("Should return pickup status list when success", func(t *testing.T) {
		repo := new(mocks.PickupStatusRepository)

		uc := usecase.NewPickupStatusUseCase(usecase.PickupStatusUsecaseImplementationConfig{
			Repository: repo,
		})

		var pickupStatus []entity.PickupStatus

		response := dto.PickupStatusLists{}

		repo.On("GetPickupStatus").Return(pickupStatus, nil)
		res, err := uc.GetPickupStatus()

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, response, *res)
	})
	t.Run("Should return error when failde get pickup status list", func(t *testing.T) {
		repo := new(mocks.PickupStatusRepository)

		uc := usecase.NewPickupStatusUseCase(usecase.PickupStatusUsecaseImplementationConfig{
			Repository: repo,
		})

		repo.On("GetPickupStatus").Return(nil, errors.New("error"))
		res, err := uc.GetPickupStatus()

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
}
