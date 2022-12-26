package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"gorm.io/gorm"
)

type PickupRepository interface {
	CreatePickup(u entity.Pickup) (*entity.Pickup, error)
	UpdateStatus(id int, status int) (*entity.Pickup, error)
}

type postgresPickupRepository struct {
	db *gorm.DB
}

type PostgresPickupRepositoryConfig struct {
	DB *gorm.DB
}

func NewPostgresPickupRepository(c PostgresPickupRepositoryConfig) PickupRepository {
	return &postgresPickupRepository{
		db: c.DB,
	}
}



func (r *postgresPickupRepository) CreatePickup(u entity.Pickup) (*entity.Pickup, error) {
	res := r.db.Create(&u)
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_PICKUP")
	}

	return &u, nil
}

func (r *postgresPickupRepository) UpdateStatus(id int, status int) (*entity.Pickup, error) {
	var u entity.Pickup
	res := r.db.Model(&u).Where("id = ?", id).Update("status_id", status)

	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_USER")
	}

	return &u, nil
}




