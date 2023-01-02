package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type WalletRepository interface {
	CreateWallet(userId int) (*entity.Wallet, error)
	IncreaseBalance(amount decimal.Decimal, wallet entity.Wallet) (*entity.Wallet, error)
	DecreaseBalance(amount decimal.Decimal, wallet entity.Wallet) (*entity.Wallet, error)
	IsValidBalance(amount decimal.Decimal, wallet entity.Wallet) bool
	GetWalletByUserID(userId int) (*entity.Wallet, error)
}

type postgresWalletRepository struct {
	db *gorm.DB
}
type PostgresWalletRepositoryConfig struct {
	DB *gorm.DB
}

func NewPostgresWalletRepository(c PostgresWalletRepositoryConfig) WalletRepository {
	return &postgresWalletRepository{
		db: c.DB,
	}
}

func (r *postgresWalletRepository) CreateWallet(userId int) (*entity.Wallet, error) {
	tx:= r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return  nil,httperror.BadRequestError(err.Error(), "ERROR_CREATING_WALLET")
	}
	u := entity.Wallet{
		Balance: decimal.NewFromInt(0),
		UserId: userId,
	}

	err := r.db.Debug().Create(&u).Error
	if err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATING_WALLET")
	}

	return &u, nil
}

func (r *postgresWalletRepository) GetWalletByUserID(userId int) (*entity.Wallet, error) {
	var u entity.Wallet


	err := r.db.Debug().Where("user_id = ?", userId).Take(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *postgresWalletRepository) IsValidBalance(amount decimal.Decimal, wallet entity.Wallet) bool {
	return wallet.Balance.GreaterThanOrEqual(amount)
}

func (r *postgresWalletRepository) IncreaseBalance(amount decimal.Decimal, wallet entity.Wallet) (*entity.Wallet, error) {
	tx:= r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return  nil,httperror.BadRequestError(err.Error(), "ERROR_UPDATING_WALLET")
	}
	wallet.Balance = wallet.Balance.Add(amount)

	err := r.db.Save(&wallet).Error
	if err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_UPDATING_WALLET")
	}

	return &wallet, nil
}

func (r *postgresWalletRepository) DecreaseBalance(amount decimal.Decimal, wallet entity.Wallet) (*entity.Wallet, error) {
	tx:= r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return  nil,httperror.BadRequestError(err.Error(), "ERROR_UPDATING_WALLET")
	}
	wallet.Balance = wallet.Balance.Sub(amount)

	err := r.db.Save(&wallet).Error
	if err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_UPDATING_WALLET")
	}

	return &wallet, nil
}
