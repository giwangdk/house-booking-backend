package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type HouseRepository interface {
	GetHouses() (*[]entity.House, error)
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

func (r *postgresHouseRepository) GetHouses() (*[]entity.House, error) {
	var cities []entity.House
	err := r.db.Find(&cities).Error
	if err != nil {
		return nil, err
	}
	return &cities, nil

}
