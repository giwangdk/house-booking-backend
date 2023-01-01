package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"
	"fmt"

	"gorm.io/gorm"
)

type ReservationRepository interface {
	CreateReservation(u entity.Reservation) (*entity.Reservation, error)
	IsHouseAvailable(checkinDate string, checkoutDate string, houseID int) (bool, error)
	GetReservationByBookingCode(code string) (*entity.Reservation, error)
	UpdateStatus(id int, status int) (*entity.Reservation, error)
	GetReservationById(id int) (*entity.Reservation, error)
	GetReservationByUserID(userId int)([]*entity.Reservation, error)
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

func (r *postgresReservationRepository) IsHouseAvailable(checkinDate string, checkoutDate string, houseID int) (bool, error) {
	var count int64

	subQuery := r.db.Select("id").Model(&entity.Reservation{}).Where("house_id = ? ", houseID)

	res := r.db.Model(&entity.Reservation{}).Where("id IN (?) AND ((check_in BETWEEN ? AND ?) OR (check_out BETWEEN ? AND ?) AND status_id != 3)", subQuery, checkinDate, checkoutDate, checkinDate, checkoutDate).Count(&count)
	if res.Error != nil {
		return false, httperror.BadRequestError(res.Error.Error(), "ERROR_CHECKING_RESERVATION")
	}

	fmt.Println(count)

	return count == 0, nil
}

func (r *postgresReservationRepository) GetReservationByBookingCode(code string) (*entity.Reservation, error) {
	var u entity.Reservation
	res := r.db.Model(entity.Reservation{}).Preload("User").Preload("House").Where("reservations.booking_code = ?", code)

	if err := res.First(&u).Error; err != nil {
		return nil, httperror.NotFoundError(err.Error())
	}
	return &u, nil
}

func (r *postgresReservationRepository) GetReservationById(id int) (*entity.Reservation, error) {
	var u entity.Reservation
	res := r.db.Model(entity.Reservation{}).Preload("User").Select("reservations.*, houses.*").Where("reservations.id = ?", id).Joins("LEFT JOIN houses ON reservations.house_id = houses.id")
	if res.First(&u).Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_GET_RESERVATION")
	}
	return &u, nil
}


func (r *postgresReservationRepository) UpdateStatus(id int, status int) (*entity.Reservation, error) {
	var u entity.Reservation
	res := r.db.Model(&u).Where("id = ?", id).Update("status_id", status)

	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_USER")
	}

	return &u, nil
}

func (r *postgresReservationRepository) GetReservationByUserID(userId int)([]*entity.Reservation, error){
	var reservations []*entity.Reservation
	res:= r.db.Model(entity.Reservation{}).Where("reservations.user_id = ? AND status_id != 3", userId).Find(&reservations)
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_GET_RESERVATION")
	}
	return reservations, nil
}

