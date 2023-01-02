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
	tx:= r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATE_HOUSE_PHOTO")
	}
	err := r.db.Create(&u).Error
	if err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATE_HOUSE_PHOTO")
	}

	return &u, nil
}

func (r *postgresHousePhotoRepository) DeleteHousePhoto(id int) error {
	tx:= r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return  httperror.BadRequestError(err.Error(), "ERROR_DELETE_HOUSE_PHOTO")
	}
	res := r.db.Delete(&entity.HousePhoto{}, id)

	if res.Error != nil {
		tx.Rollback()
		return httperror.BadRequestError(res.Error.Error(), "ERROR_DELETE_HOUSE_PHOTO")
	}

	return nil
}
