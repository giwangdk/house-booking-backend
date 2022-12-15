package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type CityRepository interface {
	GetCities() (*[]entity.City, error)
}

type postgresCityRepository struct {
	db *gorm.DB
}

type PostgresCityRepositoryConfig struct {
	DB *gorm.DB
}

func NewPostgresCityRepository(c PostgresCityRepositoryConfig) CityRepository {
	return &postgresCityRepository{
		db: c.DB,
	}
}

func (r *postgresCityRepository) GetCities() (*[]entity.City, error) {
	var cities []entity.City
	err := r.db.Find(&cities).Error
	if err != nil {
		return nil, err
	}
	return &cities, nil

}
