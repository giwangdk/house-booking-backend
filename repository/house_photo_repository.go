package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	res := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&u)

	if res.RowsAffected == 0 && res.Error == nil {
		return nil, httperror.BadRequestError("HousePhoto name already exist", "HOUSEPhoto_ALREADY_EXIST")
	}
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_HOUSEPhoto")
	}

	return &u, nil
}

func (r *postgresHousePhotoRepository) DeleteHousePhoto(id int) error {
	res := r.db.Delete(&entity.HousePhoto{}, id)

	if res.Error != nil {
		return httperror.BadRequestError(res.Error.Error(), "ERROR_DELETE_HOUSEPhoto")
	}

	return nil
}
