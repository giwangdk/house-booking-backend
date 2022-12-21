package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HouseRepository interface {
	GetHouses(page int, limit int, sortBy string, sort string, searchBy string, filterByCity int) (*[]entity.House,int, error)
	CreateHouse(u entity.House) (*entity.House, error)
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

func (r *postgresHouseRepository) GetHouses(page int, limit int, sortBy string, sort string, searchBy string, filterByCity int) (*[]entity.House,int, error) {
	var houses []entity.House
	var total int64

	subQuery:= r.db.Select("id").Table("cities").Where("name LIKE ?", "%"+searchBy+"%")

	// subQuery2 := r.db.Select("id").Table("reservations").Where("status = ?", "approved")

	res := r.db.Model(entity.House{}).Preload("City").Preload("User")
	if sortBy != "" || sort != "" {
		res = res.Order(sortBy + " " + sort)
	}
	

	if filterByCity != 0 {
		res = res.Where("city_id = ?", filterByCity)
	}

	res.Limit(limit).Offset(page)
	

	res.Where("LOWER(name) LIKE LOWER(?)", "%"+searchBy+"%").Or("city_id IN (?)", subQuery).Not("house_id").Count(&total)
	

	if err := res.Find(&houses).Error; err != nil {
		return nil,0, err
	}

	return &houses,int(total), nil
}

func (r *postgresHouseRepository) CreateHouse(u entity.House) (*entity.House, error) {
	res := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&u)

	if res.RowsAffected == 0 && res.Error == nil {
		return nil, httperror.BadRequestError("House name already exist", "HOUSE_ALREADY_EXIST")
	}
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_HOUSE")
	}

	return &u, nil
}
