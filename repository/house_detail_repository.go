package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"gorm.io/gorm"
)

type HouseDetailRepository interface {
	CreateHouseDetail(u entity.HouseDetail) (*entity.HouseDetail, error)
	GetHouseDetailById(id int) (*entity.HouseDetail, error)
	UpdateHouseDetail(u entity.HouseDetail, houseId int) (*entity.HouseDetail, error)
}

type postgresHouseDetailRepository struct {
	db *gorm.DB
}

type PostgresHouseDetailRepositoryConfig struct {
	DB *gorm.DB
}

func NewPostgresHouseDetailRepository(c PostgresHouseDetailRepositoryConfig) HouseDetailRepository {
	return &postgresHouseDetailRepository{
		db: c.DB,
	}
}

func (r *postgresHouseDetailRepository) GetHouseDetailById(id int) (*entity.HouseDetail, error) {
	var houseDetail entity.HouseDetail
	res := r.db.Model(entity.HouseDetail{}).Select("house_details.*").Where("house_details.id", id)

	if err := res.First(&houseDetail).Error; err != nil {
		return nil, httperror.NotFoundError(err.Error())
	}
	return &houseDetail, nil
}

func (r *postgresHouseDetailRepository) CreateHouseDetail(u entity.HouseDetail) (*entity.HouseDetail, error) {
	res := r.db.Create(&u)
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_HOUSE_DETAIL")
	}

	return &u, nil
}

func (r *postgresHouseDetailRepository) UpdateHouseDetail(u entity.HouseDetail, houseId int) (*entity.HouseDetail, error) {
	
	err := r.db.Where("id = ?", houseId).Updates(&u).Error

	if err != nil {
		return nil, httperror.BadRequestError(err.Error(), "ERROR_UPDATE_HOUSE_DETAIL")
	}

	return &u, nil
}
