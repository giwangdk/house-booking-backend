package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HouseDetailRepository interface {
	CreateHouseDetail(u entity.HouseDetail) (*entity.HouseDetail, error)
	GetHouseDetailById(id int) (*entity.HouseDetail, error)
	UpdateHouseDetail(u entity.HouseDetail) (*entity.HouseDetail, error)
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
	res := r.db.Model(entity.HouseDetail{}).Preload("Photos").Select("houseDetails.*, houseDetail_details.*").Where("houseDetails.id", id)

	res.Joins("LEFT JOIN houseDetail_details ON houseDetail_details.houseDetail_id = houseDetails.id")

	if err := res.First(&houseDetail).Error; err != nil {
		return nil, httperror.NotFoundError(err.Error())
	}
	return &houseDetail, nil
}

func (r *postgresHouseDetailRepository) CreateHouseDetail(u entity.HouseDetail) (*entity.HouseDetail, error) {
	res := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&u)

	if res.RowsAffected == 0 && res.Error == nil {
		return nil, httperror.BadRequestError("HouseDetail name already exist", "HOUSEDetail_ALREADY_EXIST")
	}
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_HOUSEDetail")
	}

	return &u, nil
}

func (r *postgresHouseDetailRepository) UpdateHouseDetail(u entity.HouseDetail) (*entity.HouseDetail, error) {
	err := r.db.Where("id = ?", u.ID).Updates(&u).Error

	if err != nil {
		return nil, httperror.BadRequestError(err.Error(), "ERROR_UPDATE_HOUSEDetail")
	}

	return &u, nil
}
