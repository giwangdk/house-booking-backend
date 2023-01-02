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


func TestCreateHouseDetail(t *testing.T){
	t.Run("Should return success when house detail created", func(t *testing.T) {

		repo:= new(mocks.HouseDetailRepository)

		uc:= usecase.NewHouseDetailUseCase(usecase.HouseDetailUsecaseImplementationConfig{
			Repository: repo,
		})

		req := dto.CreateHouseDetailRequest{
			MaxGuest: 		  decimal.NewFromInt(2),
			Bedrooms:            decimal.NewFromInt(2),
			Beds:                decimal.NewFromInt(2),
			Baths:               decimal.NewFromInt(2),
			HouseID:             1,
		}

		entityHouseDetail := entity.HouseDetail{
			MaxGuest: 		  req.MaxGuest,
			Bedrooms:            req.Bedrooms,
			Beds:                req.Beds,
			Baths:               req.Baths,
			HouseID:             req.HouseID,
		}

		repo.On("CreateHouseDetail", entityHouseDetail).Return(&entityHouseDetail, nil)
		houseDetail, err := uc.CreateHouseDetail(req)

		assert.Nil(t, err)
		assert.Equal(t, houseDetail.MaxGuest, req.MaxGuest)
		assert.Equal(t, houseDetail.Bedrooms, req.Bedrooms)
		assert.Equal(t, houseDetail.Beds, req.Beds)
		assert.Equal(t, houseDetail.Baths, req.Baths)
	})
	t.Run("Should return error when fail create house detail", func(t *testing.T) {

		repo:= new(mocks.HouseDetailRepository)

		uc:= usecase.NewHouseDetailUseCase(usecase.HouseDetailUsecaseImplementationConfig{
			Repository: repo,
		})

		req := dto.CreateHouseDetailRequest{
			MaxGuest: 		  decimal.NewFromInt(2),
			Bedrooms:            decimal.NewFromInt(2),
			Beds:                decimal.NewFromInt(2),
			Baths:               decimal.NewFromInt(2),
			HouseID:             1,
		}

		entityHouseDetail := entity.HouseDetail{
			MaxGuest: 		  req.MaxGuest,
			Bedrooms:            req.Bedrooms,
			Beds:                req.Beds,
			Baths:               req.Baths,
			HouseID:             req.HouseID,
		}

		repo.On("CreateHouseDetail", entityHouseDetail).Return(nil, errors.New("error"))
		houseDetail, err := uc.CreateHouseDetail(req)

		assert.NotNil(t, err)
		assert.Nil(t, houseDetail)
		assert.Equal(t, err.Error(), "error")
	})
}

func TestGetHouseDetailById(t *testing.T){
	t.Run("Should return success when get house detail", func(t *testing.T) {

		repo:= new(mocks.HouseDetailRepository)

		uc:= usecase.NewHouseDetailUseCase(usecase.HouseDetailUsecaseImplementationConfig{
			Repository: repo,
		})

		req := dto.CreateHouseDetailRequest{
			MaxGuest: 		  decimal.NewFromInt(2),
			Bedrooms:            decimal.NewFromInt(2),
			Beds:                decimal.NewFromInt(2),
			Baths:               decimal.NewFromInt(2),
			HouseID:             1,
		}

		entityHouseDetail := entity.HouseDetail{
			MaxGuest: 		  req.MaxGuest,
			Bedrooms:            req.Bedrooms,
			Beds:                req.Beds,
			Baths:               req.Baths,
			HouseID:             req.HouseID,
		}

		res:= dto.HouseDetail{
			ID:0,
			MaxGuest: 		  req.MaxGuest,
			Bedrooms:            req.Bedrooms,
			Beds:                req.Beds,
			Baths:               req.Baths,
		}

		repo.On("GetHouseDetailById", 0).Return(&entityHouseDetail, nil)
		houseDetail, err := uc.GetHouseDetailById(0)

		assert.Nil(t, err)
		assert.Equal(t, houseDetail.ID, res.ID)
		assert.Equal(t, houseDetail.MaxGuest, res.MaxGuest)
		assert.Equal(t, houseDetail.Bedrooms, res.Bedrooms)
		assert.Equal(t, houseDetail.Beds, res.Beds)

	})
	t.Run("Should return error when fail get house detail", func(t *testing.T) {

		repo:= new(mocks.HouseDetailRepository)

		uc:= usecase.NewHouseDetailUseCase(usecase.HouseDetailUsecaseImplementationConfig{
			Repository: repo,
		})


		repo.On("GetHouseDetailById", 0).Return(nil, errors.New("error"))
		houseDetail, err := uc.GetHouseDetailById(0)

		assert.NotNil(t, err)
		assert.Nil(t, houseDetail)
		assert.Equal(t, err.Error(), "error")
	})
	}

	

func TestUpdateHouseDetail(t *testing.T){
	t.Run("Should return success when success update house detail", func(t *testing.T) {
		repo:= new(mocks.HouseDetailRepository)

		uc:= usecase.NewHouseDetailUseCase(usecase.HouseDetailUsecaseImplementationConfig{
			Repository: repo,
		})

		req := dto.UpdateHouseDetailRequest{
			MaxGuest: 		  decimal.NewFromInt(2),
			Bedrooms:            decimal.NewFromInt(2),
			Beds:                decimal.NewFromInt(2),
			Baths:               decimal.NewFromInt(2),
			HouseID:             0,
		}
		entityHouseDetail := entity.HouseDetail{
			MaxGuest: 		  req.MaxGuest,
			Bedrooms:            req.Bedrooms,
			Beds:                req.Beds,
			Baths:               req.Baths,
			HouseID:             0,
		}

		res:= dto.HouseDetail{
			ID:0,
			MaxGuest: 		  req.MaxGuest,
			Bedrooms:            req.Bedrooms,
			Beds:                req.Beds,
			Baths:               req.Baths,
		}

		repo.On("GetHouseDetailById", 0).Return(&entityHouseDetail, nil)
		repo.On("UpdateHouseDetail", entityHouseDetail,0).Return(&entityHouseDetail, nil)
		houseDetail, err := uc.UpdateHouseDetail(req, 0)

		assert.Nil(t, err)
		assert.Equal(t, houseDetail.MaxGuest, res.MaxGuest)
		assert.Equal(t, houseDetail.Bedrooms, res.Bedrooms)
		assert.Equal(t, houseDetail.Beds, res.Beds)
	})
	
	t.Run("Should return error when house is not found", func(t *testing.T) {
		repo:= new(mocks.HouseDetailRepository)

		uc:= usecase.NewHouseDetailUseCase(usecase.HouseDetailUsecaseImplementationConfig{
			Repository: repo,
		})

		req := dto.UpdateHouseDetailRequest{
			MaxGuest: 		  decimal.NewFromInt(2),
			Bedrooms:            decimal.NewFromInt(2),
			Beds:                decimal.NewFromInt(2),
			Baths:               decimal.NewFromInt(2),
			HouseID:             0,
		}

		repo.On("GetHouseDetailById", 0).Return(nil, errors.New("error"))
		houseDetail, err := uc.UpdateHouseDetail(req, 0)

		assert.NotNil(t, err)
		assert.Nil(t, houseDetail)
		assert.Equal(t, err.Error(), "error")
	})
	t.Run("Should return error when fail update house detail", func(t *testing.T) {
		repo:= new(mocks.HouseDetailRepository)

		uc:= usecase.NewHouseDetailUseCase(usecase.HouseDetailUsecaseImplementationConfig{
			Repository: repo,
		})

		req := dto.UpdateHouseDetailRequest{
			MaxGuest: 		  decimal.NewFromInt(2),
			Bedrooms:            decimal.NewFromInt(2),
			Beds:                decimal.NewFromInt(2),
			Baths:               decimal.NewFromInt(2),
			HouseID:             0,
		}
		entityHouseDetail := entity.HouseDetail{
			MaxGuest: 		  req.MaxGuest,
			Bedrooms:            req.Bedrooms,
			Beds:                req.Beds,
			Baths:               req.Baths,
			HouseID:             0,
		}

		repo.On("GetHouseDetailById", 0).Return(&entityHouseDetail, nil)
		repo.On("UpdateHouseDetail", entityHouseDetail,0).Return(nil, errors.New("error"))
		houseDetail, err := uc.UpdateHouseDetail(req, 0)

		assert.NotNil(t, err)
		assert.Nil(t, houseDetail)
		assert.Equal(t, err.Error(), "error")
	})
	
	
	}

	