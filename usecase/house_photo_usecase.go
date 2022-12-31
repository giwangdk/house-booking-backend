package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/helper"
	"final-project-backend/httperror"
	"final-project-backend/repository"
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

	uploadUrl, err:= helper.ImageUploadHelper(r.Photo)
	if err != nil {
		return nil, httperror.BadRequestError("Failed to upload image", "FAILED_UPLOAD_IMAGE")
	}

	entityHousePhoto := entity.HousePhoto{
		HouseID: r.HouseID,
		Photo:   uploadUrl,
	}

	housePhoto, err := u.repository.CreateHousePhoto(entityHousePhoto)
	if err != nil {
		return nil, httperror.InternalServerError("Failed to create house photo")
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
