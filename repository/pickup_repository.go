package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"gorm.io/gorm"
)

type PickupRepository interface {
	CreatePickup(u entity.Pickup) (*entity.Pickup, error)
	GetPickups(page int, limit int, sortBy string, sort string, searchBy string, filterByStatus int) ([]entity.Pickup, int, error)
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

func (r *postgresPickupRepository) GetPickups(page int, limit int, sortBy string, sort string, searchBy string, filterByStatus int) ([]entity.Pickup, int, error) {
	var pickups []entity.Pickup
	var total int64

	
	subQuery := r.db.Select("id").Table("reservations").Where("booking_code LIKE ?", "%"+searchBy+"%")
	res := r.db.Model(entity.Pickup{})

	if filterByStatus != 0 {
		res.Where("pickup_status_id = ?", filterByStatus)
	}
		if sortBy != "" || sort != "" {
		res = res.Order(sortBy + " " + sort)
	}

	res.Preload("PickupStatus").Count(&total)
	res.Where("reservation_id IN (?)", subQuery)

	if err := res.Limit(limit).Offset(page-1).Find(&pickups).Error; err != nil {
		return nil,0, httperror.BadRequestError(err.Error(), "ERROR_GET_PICKUPS")
	}
	return pickups,int(total), nil
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
	res := r.db.Model(&u).Where("id = ?", id).Update("pickup_status_id", status).First(&u)

	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_USER")
	}

	return &u, nil
}




