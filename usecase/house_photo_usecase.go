package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
	"fmt"
)

type HousePhotoUsecase interface {
	CreateHousePhoto(r dto.CreateHousePhotoRequest) (*dto.CreateHousePhotoResponse, error)
	DeleteHousePhoto(houseId int) error
}

type HousePhotoUsecaseImplementation struct {
	repository repository.HousePhotoRepository
}

type HousePhotoUsecaseImplementationConfig struct {
	Repository repository.HousePhotoRepository
}

func NewHousePhotoUseCase(c HousePhotoUsecaseImplementationConfig) HousePhotoUsecase {
	return &HousePhotoUsecaseImplementation{
		repository: c.Repository,
	}
}

func (u *HousePhotoUsecaseImplementation) CreateHousePhoto(r dto.CreateHousePhotoRequest) (*dto.CreateHousePhotoResponse, error) {

	fmt.Println(r)

	entityHousePhoto := entity.HousePhoto{
		HouseID: r.HouseID,
		Photo:   r.Photo,
	}

	housePhoto, err := u.repository.CreateHousePhoto(entityHousePhoto)
	if err != nil {
		return nil, err
	}

	res := (&dto.CreateHousePhotoResponse{}).BuildResponse(*housePhoto)

	return res, nil

}

func (u *HousePhotoUsecaseImplementation) DeleteHousePhoto(houseId int) error {
	err := u.repository.DeleteHousePhoto(houseId)
	if err != nil {
		return err
	}

	return nil
}