package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"gorm.io/gorm"
)

type ReservationRepository interface {
	CreateReservation(u entity.Reservation) (*entity.Reservation, error)
}

type postgresReservationRepository struct {
	db *gorm.DB
}

type PostgresReservationRepositoryConfig struct {
	DB *gorm.DB
}

func NewPostgresReservationRepository(c PostgresReservationRepositoryConfig) ReservationRepository {
	return &postgresReservationRepository{
		db: c.DB,
	}
}

func (r *postgresReservationRepository) CreateReservation(u entity.Reservation) (*entity.Reservation, error) {
	res := r.db.Create(&u)
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_RESERVATION")
	}

	return &u, nil
}


