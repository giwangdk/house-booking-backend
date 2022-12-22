package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"gorm.io/gorm"
)

type HousePhotoRepository interface {
	CreateHousePhoto(u entity.HousePhoto) (*entity.HousePhoto, error)
	DeleteHousePhoto(id int) error
}

type postgresHousePhotoRepository struct {
	db *gorm.DB
}

type PostgresHousePhotoRepositoryConfig struct {
	DB *gorm.DB
}

func NewPostgresHousePhotoRepository(c PostgresHousePhotoRepositoryConfig) HousePhotoRepository {
	return &postgresHousePhotoRepository{
		db: c.DB,
	}
}

func (r *postgresHousePhotoRepository) CreateHousePhoto(u entity.HousePhoto) (*entity.HousePhoto, error) {
	err := r.db.Create(&u).Error

	if err != nil {
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATE_HOUSE_PHOTO")
	}

	return &u, nil
}

func (r *postgresHousePhotoRepository) DeleteHousePhoto(id int) error {
	res := r.db.Delete(&entity.HousePhoto{}, id)

	if res.Error != nil {
		return httperror.BadRequestError(res.Error.Error(), "ERROR_DELETE_HOUSE_PHOTO")
	}

	return nil
}
