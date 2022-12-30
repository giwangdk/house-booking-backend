package repository

import (
	"final-project-backend/entity"
	"final-project-backend/helper"
	"final-project-backend/httperror"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HouseRepository interface {
	GetHouses(userId int, page int, limit int, sortBy string, sort string, searchBy string, filterByCity int, checkIn string, checkOut string) (*[]entity.House, int, error)
	CreateHouse(u entity.HouseProfile) (*entity.HouseProfile, error)
	GetHouseById(id int) (*entity.House, error)
	UpdateHouse(u entity.HouseProfile, userId int) (*entity.HouseProfile, error)
	DeleteHouse(id int) error
}

type postgresHouseRepository struct {
	db *gorm.DB
}

type PostgresHouseRepositoryConfig struct {
	DB *gorm.DB
}

func NewPostgresHouseRepository(c PostgresHouseRepositoryConfig) HouseRepository {
	return &postgresHouseRepository{
		db: c.DB,
	}
}

func (r *postgresHouseRepository) GetHouses(userId int, page int, limit int, sortBy string, sort string, searchBy string, filterByCity int, checkIn string, checkOut string) (*[]entity.House, int, error) {
	var houses []entity.House

	var total int64

	subQuery := r.db.Debug().Select("id").Table("cities").Where("LOWER(name) LIKE lOWER(?)", "%"+searchBy+"%")

	subQuery2 := r.db.Debug().Select("house_id").Table("reservations").Where("check_in between ? and ? or check_out between ? and ? and status_id != 3", checkIn, checkOut, checkIn, checkOut)

	res := r.db.Model(entity.House{})

	res.Preload("Photos").Preload("City").Preload("User").Select("houses.*, house_details.*")

	if sortBy != "" && sort != "" {

		if sortBy == "city"{
			res.Order("cities.name " + sort)
		}
		res.Order(sortBy + " " + sort)
	}

	if filterByCity != 0 {
		res = res.Where("city_id = ?", filterByCity)
	}

	if searchBy != "" {
		res.Where("LOWER(name) LIKE LOWER(?)", "%"+searchBy+"%").Or("city_id IN (?)", subQuery).Count(&total)

	}
	if checkIn != "" && checkOut != "" {
		res.Where("houses.id NOT IN (?)", subQuery2)
	}

	if userId != 0 {
		res = res.Where("houses.user_id = ?", userId)
	}
	res.Count(&total)
	res.Joins("LEFT JOIN house_details ON house_details.house_id = houses.id")
	res.Scopes(helper.Paginate(page, limit),)

	if err := res.Find(&houses).Error; err != nil {
		return nil, 0, err
	}

	return &houses, int(total), nil
}

func (r *postgresHouseRepository) GetHouseById(id int) (*entity.House, error) {
	var house entity.House
	res := r.db.Model(entity.House{}).Preload("Photos").Preload("City").Select("houses.*, house_details.*").Where("houses.id", id)

	res.Joins("LEFT JOIN house_details ON house_details.house_id = houses.id")

	if err := res.First(&house).Error; err != nil {
		return nil, httperror.NotFoundError(err.Error())
	}
	return &house, nil
}

func (r *postgresHouseRepository) CreateHouse(u entity.HouseProfile) (*entity.HouseProfile, error) {
	res := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&u)

	if res.RowsAffected == 0 && res.Error == nil {
		return nil, httperror.BadRequestError("House Name already Exist!", "HOUSE_ALREADY_EXIST")
	}
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_HOUSE")
	}

	return &u, nil
}

func (r *postgresHouseRepository) UpdateHouse(u entity.HouseProfile, userId int) (*entity.HouseProfile, error) {
	err := r.db.Where("id = ?", userId).Updates(&u).Error

	if err != nil {
		return nil, httperror.BadRequestError(err.Error(), "ERROR_UPDATE_HOUSE")
	}

	return &u, nil
}


func (r *postgresHouseRepository) DeleteHouse(id int) error {
	err := r.db.Where("id = ?", id).Delete(&entity.House{}).Error
	if err != nil {
		return httperror.BadRequestError(err.Error(), "ERROR_DELETE_HOUSE")
	}

	err= r.db.Where("house_id = ?", id).Delete(&entity.HouseDetail{}).Error
	if err != nil {
		return httperror.BadRequestError(err.Error(), "ERROR_DELETE_HOUSE")
	}

	err= r.db.Where("house_id = ?", id).Delete(&entity.HousePhoto{}).Error

	if err != nil {
		return httperror.BadRequestError(err.Error(), "ERROR_DELETE_HOUSE")
	}

	return nil
}