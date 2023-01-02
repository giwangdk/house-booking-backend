package usecase_test

import (
	"errors"
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"final-project-backend/usecase"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)


func TestGetHouses(t *testing.T){
	t.Run("Should return success when get houses", func(t *testing.T) {
		houseRepo:= new(mocks.HouseRepository)
		reservationRepo:= new(mocks.ReservationRepository)

		uc:= usecase.NewHouseUseCase(usecase.HouseUsecaseImplementationConfig{
			Repository: houseRepo,
			ReservationRepo : reservationRepo,
		})

		houses:= []entity.House{
			{
				Name: "House 1",
				Price: decimal.NewFromInt(1000000),
			},
		}

		houseRepo.On("GetHouses" ,0, 0,0,"","","",0 , "", "").Return(&houses,1, nil)
		_, err := uc.GetHouses(0,0,"","","",0 , "", "")

		assert.Nil(t, err)
		assert.Equal(t, houses[0].Name, "House 1")	
	})
	t.Run("Should return error when fail to get houses", func(t *testing.T) {
		houseRepo:= new(mocks.HouseRepository)
		reservationRepo:= new(mocks.ReservationRepository)

		uc:= usecase.NewHouseUseCase(usecase.HouseUsecaseImplementationConfig{
			Repository: houseRepo,
			ReservationRepo : reservationRepo,
		})

		houseRepo.On("GetHouses" ,0, 0,0,"","","",0 , "", "").Return(nil,0, errors.New("error"))
		res, err := uc.GetHouses(0,0,"","","",0 , "", "")

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, err.Error(), "error")
	})
}

func TestCreateHouse(t *testing.T){
	t.Run("Should return success when house created", func(t *testing.T) {
		houseRepo:= new(mocks.HouseRepository)
		reservationRepo:= new(mocks.ReservationRepository)

		uc:= usecase.NewHouseUseCase(usecase.HouseUsecaseImplementationConfig{
			Repository: houseRepo,
			ReservationRepo : reservationRepo,
		})

		houseProfile := entity.HouseProfile{
			Name: "House 1",
		}

		req:= dto.CreateHouseRequest{
			Name: "House 1",
		}

		houseRepo.On("CreateHouse" ,houseProfile).Return(&houseProfile, nil)
		res, err := uc.CreateHouse(req)

		assert.Nil(t, err)
		assert.Equal(t, res.Name, "House 1")
	})
	t.Run("Should return error when fail to create house", func(t *testing.T) {
		houseRepo:= new(mocks.HouseRepository)
		reservationRepo:= new(mocks.ReservationRepository)

		uc:= usecase.NewHouseUseCase(usecase.HouseUsecaseImplementationConfig{
			Repository: houseRepo,
			ReservationRepo : reservationRepo,
		})
		req:= dto.CreateHouseRequest{
			Name: "House 1",
		}

		houseProfile := entity.HouseProfile{
			Name: "House 1",
		}

		houseRepo.On("CreateHouse" ,houseProfile).Return(nil, errors.New("error"))
		res, err := uc.CreateHouse(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, err.Error(), "Failed to create house!")
	})
}

func TestUpdateHouse(t *testing.T){
	t.Run("Should return success when house updated", func(t *testing.T) {
		houseRepo:= new(mocks.HouseRepository)
		reservationRepo:= new(mocks.ReservationRepository)

		uc:= usecase.NewHouseUseCase(usecase.HouseUsecaseImplementationConfig{
			Repository: houseRepo,
			ReservationRepo : reservationRepo,
		})

		house:= entity.House{
			Name: "House 1",
			Price: decimal.NewFromInt(1000000),
		}

		houseProfile := entity.HouseProfile{
			Name: "House 1",
		}

		req:= dto.UpdateHouseRequest{
			Name: "House 1",
		}
		houseRepo.On("GetHouseById" ,0).Return(&house, nil)
		houseRepo.On("UpdateHouse" ,houseProfile, 0).Return(&houseProfile, nil)
		res, err := uc.UpdateHouse(req, 0)

		assert.Nil(t, err)
		assert.Equal(t, res.Name, "House 1")
	})
	t.Run("Should return error when fail get house by id", func(t *testing.T) {
		houseRepo:= new(mocks.HouseRepository)
		reservationRepo:= new(mocks.ReservationRepository)

		uc:= usecase.NewHouseUseCase(usecase.HouseUsecaseImplementationConfig{
			Repository: houseRepo,
			ReservationRepo : reservationRepo,
		})


		req:= dto.UpdateHouseRequest{
			Name: "House 1",
		}
		houseRepo.On("GetHouseById" ,0).Return(nil, errors.New("error"))
		res, err := uc.UpdateHouse(req, 0)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
	t.Run("Should return error when fail update house", func(t *testing.T) {
		houseRepo:= new(mocks.HouseRepository)
		reservationRepo:= new(mocks.ReservationRepository)

		uc:= usecase.NewHouseUseCase(usecase.HouseUsecaseImplementationConfig{
			Repository: houseRepo,
			ReservationRepo : reservationRepo,
		})

		house:= entity.House{
			Name: "House 1",
			Price: decimal.NewFromInt(1000000),
		}

		houseProfile := entity.HouseProfile{
			Name: "House 1",
		}

		req:= dto.UpdateHouseRequest{
			Name: "House 1",
		}
		houseRepo.On("GetHouseById" ,0).Return(&house, nil)
		houseRepo.On("UpdateHouse" ,houseProfile, 0).Return(nil, errors.New("error"))
		res, err := uc.UpdateHouse(req, 0)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}

func TestGetHousesHost(t *testing.T){
	t.Run("Should return success when get houses host", func(t *testing.T) {
		houseRepo:= new(mocks.HouseRepository)
		reservationRepo:= new(mocks.ReservationRepository)

		uc:= usecase.NewHouseUseCase(usecase.HouseUsecaseImplementationConfig{
			Repository: houseRepo,
			ReservationRepo : reservationRepo,
		})
		
		var houses []entity.House

		houseRepo.On("GetHouses" ,0,0,0,"","","",0 , "", "").Return(&houses,0, nil)
		_, err := uc.GetHousesHost(0,0,0,"","","")

		assert.Nil(t, err)
		
	})
	t.Run("Should return error when failed get houses host", func(t *testing.T) {
		houseRepo:= new(mocks.HouseRepository)
		reservationRepo:= new(mocks.ReservationRepository)

		uc:= usecase.NewHouseUseCase(usecase.HouseUsecaseImplementationConfig{
			Repository: houseRepo,
			ReservationRepo : reservationRepo,
		})
		

		houseRepo.On("GetHouses" ,0,0,0,"","","",0 , "", "").Return(nil,0, errors.New("error"))
		res, err := uc.GetHousesHost(0,0,0,"","","")

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, err.Error(), "error")
		
	})
}

// func TestDeleteHouse(t *testing.T){
// 	t.Run("Should return success when house deleted", func(t *testing.T) {
// 		houseRepo:= new(mocks.HouseRepository)
// 		reservationRepo:= new(mocks.ReservationRepository)

// 		uc:= usecase.NewHouseUseCase(usecase.HouseUsecaseImplementationConfig{
// 			Repository: houseRepo,
// 			ReservationRepo : reservationRepo,
// 		})

// 		house:= entity.House{
// 			Name: "House 1",
// 			Price: decimal.NewFromInt(1000000),
// 		}

// 		houseRepo.On("GetHouseById" ,0).Return(&house, nil)
// 		houseRepo.On("DeleteHouse" ,0).Return(nil)
// 		_, err := uc.DeleteHouse(0)

// 		assert.Nil(t, err)
// })
// }