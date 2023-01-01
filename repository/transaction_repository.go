package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(u entity.Transaction) (*entity.Transaction, error)
	GetTransactionByBookingCode(bookingCode string) (*entity.Transaction, error)
	GetTransactionsGuest() ([]entity.Transaction, error)
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

func (r * postgresTransactionRepository) GetTransactionsGuest() ([]entity.Transaction, error) {
	subQuery := r.db.Table("reservations").Select("id").Where("status_id = ?", 4)
	var transactions []entity.Transaction

	res := r.db.Where("reservation_id IN (?)", subQuery).Preload("Reservation").Find(&transactions)

	if res.Error != nil {
		return nil, httperror.NotFoundError(res.Error.Error())
	}

	return transactions, nil
}


func (r *postgresTransactionRepository) GetTransactionByBookingCode(bookingCode string) (*entity.Transaction, error) {
	var transaction entity.Transaction
	subQuery := r.db.Table("reservations").Select("id").Where("booking_code = ?", bookingCode)

	res:= r.db.Where("reservation_id IN (?)", subQuery).First(&transaction)
	if res.Error != nil {
		return nil, httperror.NotFoundError(res.Error.Error())
	}

	return &transaction, nil
}



