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

func TestGetPickups(t *testing.T) {
	t.Run("Should return pickup list when success", func(t *testing.T) {
		repo := new(mocks.PickupRepository)
		uc := usecase.NewPickupUseCase(usecase.PickupUsecaseImplementationConfig{
			Repository: repo,
		})

		pickupList := dto.PickupLists{
			Pickups: []dto.Pickup{
				{
					ReservationID: 1,
				},
			},
		}

		var pickups = []entity.Pickup{
			{
				ReservationID: 1,
			},
		}

		repo.On("GetPickups", 0, 0, "id", "asc", "", 0).Return(pickups, 0, nil)
		res, err := uc.GetPickups(0, 0, "id", "asc", "", 0)

		assert.Nil(t, err)
		assert.Equal(t, pickupList, *res)
		assert.NotNil(t, res)
	})
	t.Run("Should return error when fail get pickup list", func(t *testing.T) {
		repo := new(mocks.PickupRepository)
		uc := usecase.NewPickupUseCase(usecase.PickupUsecaseImplementationConfig{
			Repository: repo,
		})
		repo.On("GetPickups", 0, 0, "id", "asc", "", 0).Return(nil, 0, errors.New("error"))
		res, err := uc.GetPickups(0, 0, "id", "asc", "", 0)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
}

func TestCreatePickup(t *testing.T) {
	t.Run("Should return pickup when success", func(t *testing.T) {
		repo := new(mocks.PickupRepository)
		uc := usecase.NewPickupUseCase(usecase.PickupUsecaseImplementationConfig{
			Repository: repo,
		})

		req := dto.CreatePickupRequest{
			ReservationID: 1,
		}

		pickup := entity.Pickup{
			ReservationID:  1,
			PickupStatusID: 1,
		}

		repo.On("CreatePickup", pickup).Return(&pickup, nil)
		res, err := uc.CreatePickup(req)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, req.ReservationID, res.ReservationID)
	})
	t.Run("Should return error when fail create pickup", func(t *testing.T) {
		repo := new(mocks.PickupRepository)
		uc := usecase.NewPickupUseCase(usecase.PickupUsecaseImplementationConfig{
			Repository: repo,
		})

		req := dto.CreatePickupRequest{
			ReservationID: 1,
		}

		pickup := entity.Pickup{
			ReservationID:  1,
			PickupStatusID: 1,
		}

		repo.On("CreatePickup", pickup).Return(nil, errors.New("error"))
		res, err := uc.CreatePickup(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Failed to create pickup", err.Error())
	})
}

func TestUpdateStatusPickup(t *testing.T) {
	t.Run("Should return pickup when success update", func(t *testing.T) {
		repo := new(mocks.PickupRepository)
		uc := usecase.NewPickupUseCase(usecase.PickupUsecaseImplementationConfig{
			Repository: repo,
		})

		pickup := entity.Pickup{
			PickupStatusID: 2,
		}

		repo.On("UpdateStatus", 0, 2).Return(&pickup, nil)
		res, err := uc.UpdateStatusPickup(0, 2)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 0, res.UserID)
	})
	t.Run("Should return error when failed update status pickup", func(t *testing.T) {
		repo := new(mocks.PickupRepository)
		uc := usecase.NewPickupUseCase(usecase.PickupUsecaseImplementationConfig{
			Repository: repo,
		})

		repo.On("UpdateStatus", 0, 2).Return(nil, errors.New("error"))
		res, err := uc.UpdateStatusPickup(0, 2)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Failed to update status", err.Error())
	})
}
