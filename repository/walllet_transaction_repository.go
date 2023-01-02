package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"
	"fmt"

	"gorm.io/gorm"
)

type WalletTransactionRepository interface {
	CreateWalletTransaction(u entity.WalletTransaction) (*entity.WalletTransaction, error)
	GetWalletTransactionsUser(walletID int) (*[]entity.WalletTransaction, error)
	GetWalletTransactions(s string, sortBy string, sort string, limit int, page int) (*[]entity.WalletTransaction, error)
}

type postgresWalletTransactionRepository struct {
	db *gorm.DB
}
type PostgresWalletTransactionRepositoryConfig struct {
	DB *gorm.DB
}

func NewPostgresWalletTransactionRepository(c PostgresWalletTransactionRepositoryConfig) WalletTransactionRepository {
	return &postgresWalletTransactionRepository{
		db: c.DB,
	}
}

func (r *postgresWalletTransactionRepository) CreateWalletTransaction(u entity.WalletTransaction) (*entity.WalletTransaction, error) {
	tx:= r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return  nil,httperror.BadRequestError(err.Error(), "ERROR_CREATE_WALLET_TRANSACTION")
	}
	err := r.db.Create(&u).Error

	if err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATE_WALLET_TRANSACTION")
	}

	return &u, nil
}

func (r *postgresWalletTransactionRepository) GetWalletTransactionsUser(walletID int) (*[]entity.WalletTransaction, error) {
	var Wallettransactions []entity.WalletTransaction

	fmt.Println("wallet", walletID)

	err := r.db.Where("sender = ? OR recipient = ?", walletID, walletID).Order("created_at desc").Limit(10).Find(&Wallettransactions).Error

	if err != nil {
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATE_WALLET_TRANSACTION")
	}

	return &Wallettransactions, nil
}

func (r *postgresWalletTransactionRepository) GetWalletTransactions(s string, sortBy string, sort string, limit int, page int) (*[]entity.WalletTransaction, error) {
	var Wallettransactions []entity.WalletTransaction

	err := r.db.Debug().Where("description LIKE ?", "%"+s+"%").Order(sortBy + " " + sort).Limit(limit).Offset(page).Find(&Wallettransactions).Error

	if err != nil {
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATE_WALLET_TRANSACTION")
	}

	return &Wallettransactions, nil
}


