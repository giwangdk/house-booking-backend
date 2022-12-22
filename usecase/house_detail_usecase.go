package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
)

type HouseDetailUsecase interface {
	CreateHouseDetail(r dto.CreateHouseDetailRequest) (*dto.CreateHouseDetailResponse, error)
	GetHouseDetailById(houseDetailId int) (*dto.HouseDetail, error)
	UpdateHouseDetail(r dto.UpdateHouseDetailRequest, houseDetailId int) (*dto.UpdateHouseDetailResponse, error)
}

type HouseDetailUsecaseImplementation struct {
	repository repository.HouseDetailRepository
}

type HouseDetailUsecaseImplementationConfig struct {
	Repository repository.HouseDetailRepository
}

func NewHouseDetailUseCase(c HouseDetailUsecaseImplementationConfig) HouseDetailUsecase {
	return &HouseDetailUsecaseImplementation{
		repository: c.Repository,
	}
}

func (u *HouseDetailUsecaseImplementation) CreateHouseDetail(r dto.CreateHouseDetailRequest) (*dto.CreateHouseDetailResponse, error) {

	entityHouseDetail := entity.HouseDetail{
		Bedrooms:            r.Bedrooms,
		Beds:                r.Beds,
		Baths:               r.Baths,
		HouseFacilities:     r.HouseFacilities,
		HouseRules:          r.HouseRules,
		HouseServices:       r.HouseServices,
		BathroomsFacilities: r.BathroomsFacilities,
		HouseID:             r.HouseID,
	}

	houseDetail, err := u.repository.CreateHouseDetail(entityHouseDetail)
	if err != nil {
		return nil, err
	}

	res := (&dto.CreateHouseDetailResponse{}).BuildResponse(*houseDetail)

	return res, nil

}

func (u *HouseDetailUsecaseImplementation) GetHouseDetailById(houseDetailId int) (*dto.HouseDetail, error) {
	houseDetail, err := u.repository.GetHouseDetailById(houseDetailId)
	if err != nil {
		return nil, err
	}

	res := (&dto.HouseDetail{}).BuildResponse(*houseDetail)

	return res, nil
}

func (u *HouseDetailUsecaseImplementation) UpdateHouseDetail(r dto.UpdateHouseDetailRequest, houseDetailId int) (*dto.UpdateHouseDetailResponse, error) {
	_, err := u.GetHouseDetailById(houseDetailId)
	if err != nil {
		return nil, err
	}

	entity := entity.HouseDetail{
		MaxGuest:            r.MaxGuest,
		Bedrooms:            r.Bedrooms,
		Beds:                r.Beds,
		Baths:               r.Baths,
		HouseFacilities:     r.HouseFacilities,
		HouseRules:          r.HouseRules,
		HouseServices:       r.HouseServices,
		BathroomsFacilities: r.BathroomsFacilities,
	}

	updatedUser, err := u.repository.UpdateHouseDetail(entity)
	if err != nil {
		return nil, err
	}

	res := (&dto.UpdateHouseDetailResponse{}).BuildResponse(*updatedUser)
	return res, nil
}
