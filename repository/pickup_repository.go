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
	subQuery2 := r.db.Select("id").Table("reservations").Where("status_id = 2")
	res := r.db.Model(entity.Pickup{})

	if filterByStatus != 0 {
		res.Where("pickup_status_id = ?", filterByStatus)
	}
		if sortBy != "" || sort != "" {
		res = res.Order(sortBy + " " + sort)
	}


	res.Preload("PickupStatus").Preload("Reservation")
	res.Where("reservation_id IN (?)", subQuery2)
	res.Where("reservation_id IN (?)", subQuery).Count(&total)
	res.Limit(limit).Offset(page-1)
	if err := res.Find(&pickups).Error; err != nil {
		return nil,0, httperror.BadRequestError(err.Error(), "ERROR_GET_PICKUPS")
	}
	return pickups,int(total), nil
}

func (r *postgresPickupRepository) CreatePickup(u entity.Pickup) (*entity.Pickup, error) {
	tx:= r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATE_PICKUP")
	}
	res := r.db.Create(&u)
	if res.Error != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_PICKUP")
	}

	return &u, nil
}

func (r *postgresPickupRepository) UpdateStatus(id int, status int) (*entity.Pickup, error) {
	var u entity.Pickup
	tx:= r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_UPDATE_PICKUP")
	}
	res := r.db.Model(&u).Where("id = ?", id).Update("pickup_status_id", status).First(&u)

	if res.Error != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_UPDATE_PICKUP")
	}

	return &u, nil
}




