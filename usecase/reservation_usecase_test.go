package usecase_test

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"final-project-backend/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateReservation(t *testing.T) {
	t.Run("Should return success when reservation is  created", func(t *testing.T) {

		repo := new(mocks.ReservationRepository)
		userUsecase := new(mocks.UserUsecase)
		pickupUsecase := new(mocks.PickupUsecase)

		uc := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
			Repository:    repo,
			UserUsecase:   userUsecase,
			PickupUsecase: pickupUsecase,
		})

		entityReservation := entity.Reservation{
			CheckIn:  "2021-01-01",
			CheckOut: "2021-01-02",
			HouseID:  1,
		}

		repo.On("IsHouseAvailable", "2021-01-01", "2021-01-02", 1).Return(true, nil)
		repo.On("CreateReservation", entityReservation).Return(&entityReservation, nil)
		reservation, err := uc.CreateReservation(entityReservation)

		assert.Nil(t, err)
		assert.NotNil(t, reservation)
	})
	t.Run("Should return error when house is not available", func(t *testing.T) {

		repo := new(mocks.ReservationRepository)
		userUsecase := new(mocks.UserUsecase)
		pickupUsecase := new(mocks.PickupUsecase)

		uc := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
			Repository:    repo,
			UserUsecase:   userUsecase,
			PickupUsecase: pickupUsecase,
		})

		entityReservation := entity.Reservation{
			CheckIn:  "2021-01-01",
			CheckOut: "2021-01-02",
			HouseID:  1,
		}

		repo.On("IsHouseAvailable", "2021-01-01", "2021-01-02", 1).Return(false, nil)
		reservation, err := uc.CreateReservation(entityReservation)

		assert.NotNil(t, err)
		assert.Nil(t, reservation)
		assert.Equal(t, "House is not available", err.Error())
	})
	t.Run("Should return error when house is not available", func(t *testing.T) {

		repo := new(mocks.ReservationRepository)
		userUsecase := new(mocks.UserUsecase)
		pickupUsecase := new(mocks.PickupUsecase)

		uc := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
			Repository:    repo,
			UserUsecase:   userUsecase,
			PickupUsecase: pickupUsecase,
		})

		entityReservation := entity.Reservation{
			CheckIn:  "2021-01-01",
			CheckOut: "2021-01-02",
			HouseID:  1,
		}

		repo.On("IsHouseAvailable", "2021-01-01", "2021-01-02", 1).Return(true, errors.New("error"))
		reservation, err := uc.CreateReservation(entityReservation)

		assert.NotNil(t, err)
		assert.Nil(t, reservation)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("Should return error when reservation is fail", func(t *testing.T) {

		repo := new(mocks.ReservationRepository)
		userUsecase := new(mocks.UserUsecase)
		pickupUsecase := new(mocks.PickupUsecase)

		uc := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
			Repository:    repo,
			UserUsecase:   userUsecase,
			PickupUsecase: pickupUsecase,
		})

		entityReservation := entity.Reservation{
			CheckIn:  "2021-01-01",
			CheckOut: "2021-01-02",
			HouseID:  1,
		}

		repo.On("IsHouseAvailable", "2021-01-01", "2021-01-02", 1).Return(true, nil)
		repo.On("CreateReservation", entityReservation).Return(nil, errors.New("error"))
		reservation, err := uc.CreateReservation(entityReservation)

		assert.NotNil(t, err)
		assert.Nil(t, reservation)
		assert.Equal(t, "error", err.Error())
	})
}

func TestGetReservationByBookingCode(t *testing.T) {
	t.Run("Should return success when reservation is  created", func(t *testing.T) {

		repo := new(mocks.ReservationRepository)
		userUsecase := new(mocks.UserUsecase)
		pickupUsecase := new(mocks.PickupUsecase)

		uc := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
			Repository:    repo,
			UserUsecase:   userUsecase,
			PickupUsecase: pickupUsecase,
		})

		entityReservation := entity.Reservation{
			CheckIn:  "2021-01-01",
			CheckOut: "2021-01-02",
			HouseID:  1,
		}

		repo.On("GetReservationByBookingCode", "123").Return(&entityReservation, nil)
		reservation, err := uc.GetReservationByBookingCode("123")

		assert.Nil(t, err)
		assert.NotNil(t, reservation)
	})
	t.Run("Should return success when reservation is  created", func(t *testing.T) {

		repo := new(mocks.ReservationRepository)
		userUsecase := new(mocks.UserUsecase)
		pickupUsecase := new(mocks.PickupUsecase)

		uc := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
			Repository:    repo,
			UserUsecase:   userUsecase,
			PickupUsecase: pickupUsecase,
		})

		repo.On("GetReservationByBookingCode", "123").Return(nil, errors.New("error"))
		reservation, err := uc.GetReservationByBookingCode("123")

		assert.NotNil(t, err)
		assert.Nil(t, reservation)
		assert.Equal(t, "Reservation not found", err.Error())
	})
}

func TestGetReservationById(t *testing.T) {
	t.Run("Should return reservation when the success ", func(t *testing.T) {

		repo := new(mocks.ReservationRepository)
		userUsecase := new(mocks.UserUsecase)
		pickupUsecase := new(mocks.PickupUsecase)

		uc := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
			Repository:    repo,
			UserUsecase:   userUsecase,
			PickupUsecase: pickupUsecase,
		})

		entityReservation := entity.Reservation{
			CheckIn:  "2021-01-01",
			CheckOut: "2021-01-02",
			HouseID:  1,
		}

		repo.On("GetReservationById", 1).Return(&entityReservation, nil)
		reservation, err := uc.GetReservationById(1)

		assert.Nil(t, err)
		assert.NotNil(t, reservation)
	})
	t.Run("Should return error when reservation not foung", func(t *testing.T) {

		repo := new(mocks.ReservationRepository)
		userUsecase := new(mocks.UserUsecase)
		pickupUsecase := new(mocks.PickupUsecase)

		uc := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
			Repository:    repo,
			UserUsecase:   userUsecase,
			PickupUsecase: pickupUsecase,
		})

		repo.On("GetReservationById", 1).Return(nil, errors.New("error"))
		reservation, err := uc.GetReservationById(1)

		assert.NotNil(t, err)
		assert.Nil(t, reservation)
		assert.Equal(t, "Reservation not found", err.Error())
	})
}

func TestUpdateStatusReservation(t *testing.T) {
	t.Run("Should return success when status reservation is updated", func(t *testing.T) {

		repo := new(mocks.ReservationRepository)
		userUsecase := new(mocks.UserUsecase)
		pickupUsecase := new(mocks.PickupUsecase)

		uc := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
			Repository:    repo,
			UserUsecase:   userUsecase,
			PickupUsecase: pickupUsecase,
		})

		entityReservation := entity.Reservation{
			CheckIn:  "2021-01-01",
			CheckOut: "2021-01-02",
			HouseID:  1,
			StatusID: 2,
		}

		repo.On("UpdateStatus", 1, 2).Return(&entityReservation, nil)
		reservation, err := uc.UpdateStatusReservation(1, 2)

		assert.Nil(t, err)
		assert.NotNil(t, &reservation)
		assert.Equal(t, entityReservation.StatusID, reservation.StatusID)
	})
	t.Run("Should return success when status reservation is updated", func(t *testing.T) {

		repo := new(mocks.ReservationRepository)
		userUsecase := new(mocks.UserUsecase)
		pickupUsecase := new(mocks.PickupUsecase)

		uc := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
			Repository:    repo,
			UserUsecase:   userUsecase,
			PickupUsecase: pickupUsecase,
		})

		repo.On("UpdateStatus", 1, 2).Return(nil, errors.New("error"))
		reservation, err := uc.UpdateStatusReservation(1, 2)

		assert.NotNil(t, err)
		assert.Nil(t, reservation)
		assert.Equal(t, "Failed to update status", err.Error())
	})
}

func TestGetReservationByUserId(t *testing.T) {
	t.Run("Should return reservation when the success ", func(t *testing.T) {

		repo := new(mocks.ReservationRepository)
		userUsecase := new(mocks.UserUsecase)
		pickupUsecase := new(mocks.PickupUsecase)

		uc := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
			Repository:    repo,
			UserUsecase:   userUsecase,
			PickupUsecase: pickupUsecase,
		})

		var reservations []*entity.Reservation

		repo.On("GetReservationByUserID", 1).Return(reservations, nil)
		reservation, err := uc.GetReservationByUserId(1)

		assert.Nil(t, err)
		assert.NotNil(t, reservation)
	})
	t.Run("Should return error when reservation not found", func(t *testing.T) {

		repo := new(mocks.ReservationRepository)
		userUsecase := new(mocks.UserUsecase)
		pickupUsecase := new(mocks.PickupUsecase)

		uc := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
			Repository:    repo,
			UserUsecase:   userUsecase,
			PickupUsecase: pickupUsecase,
		})

		repo.On("GetReservationByUserID", 1).Return(nil, errors.New("error"))
		reservation, err := uc.GetReservationByUserId(1)

		assert.NotNil(t, err)
		assert.Nil(t, reservation)
		assert.Equal(t, "Reservation not found", err.Error())
	})
}
