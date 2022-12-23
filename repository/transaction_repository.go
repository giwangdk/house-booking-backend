package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(u entity.Transaction) (*entity.Transaction, error)
}

type postgresTransactionRepository struct {
	db *gorm.DB
}

type PostgresTransactionRepositoryConfig struct {
	DB *gorm.DB
}

func NewPostgresTransactionRepository(c PostgresTransactionRepositoryConfig) TransactionRepository {
	return &postgresTransactionRepository{
		db: c.DB,
	}
}



func (r *postgresTransactionRepository) CreateTransaction(u entity.Transaction) (*entity.Transaction, error) {
	res := r.db.Create(&u)
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_TRANSACTION")
	}

	return &u, nil
}




