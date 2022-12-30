package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"gorm.io/gorm"
)

type PickupStatusRepository interface {
	GetPickupStatus() ([]entity.PickupStatus, error)
}

type postgresPickupStatusRepository struct {
	db *gorm.DB
}

type PostgresPickupStatusRepositoryConfig struct {
	DB *gorm.DB
}

func NewPostgresPickupStatusRepository(c PostgresPickupStatusRepositoryConfig) PickupStatusRepository {
	return &postgresPickupRepository{
		db: c.DB,
	}
}

func (r *postgresPickupRepository) GetPickupStatus() ([]entity.PickupStatus, error) {
	var pickupStatus []entity.PickupStatus
	res:= r.db.Find(&pickupStatus)

	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_GET_PICKUPS_STATUS")
	}
	return pickupStatus, nil
}





