package usecase

import (
	"final-project-backend/entity"
	"final-project-backend/repository"

	"github.com/shopspring/decimal"
)

type WalletUsecase interface {
	CreateWallet(userID int) (*entity.Wallet, error)
	IncreaseBalance(amount decimal.Decimal, wallet entity.Wallet) (*entity.Wallet, error)
	DecreaseBalance(amount decimal.Decimal, wallet entity.Wallet) (*entity.Wallet, error)
	GetWalletByUserID(userId int) (*entity.Wallet, error)
	IsValidBalance(amount decimal.Decimal, wallet entity.Wallet) bool
}

type walletUsecaseImplementation struct {
	repository repository.WalletRepository
}

type WalletUsecaseImplementationConfig struct {
	Repository repository.WalletRepository
}

func NewWalletUseCase(c WalletUsecaseImplementationConfig) WalletUsecase {
	return &walletUsecaseImplementation{
		repository: c.Repository,
	}
}

func (u *walletUsecaseImplementation) CreateWallet(userId int) (*entity.Wallet, error) {

	w, err := u.repository.CreateWallet(userId)
	if err != nil {
		return nil, err
	}

	return w, nil
}
func (u *walletUsecaseImplementation) IsValidBalance(amount decimal.Decimal, wallet entity.Wallet) bool {
	return u.repository.IsValidBalance(amount, wallet)
}

func (u *walletUsecaseImplementation) GetWalletByUserID(userId int) (*entity.Wallet, error) {
	w, err := u.repository.GetWalletByUserID(userId)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (u *walletUsecaseImplementation) IncreaseBalance(amount decimal.Decimal, wallet entity.Wallet) (*entity.Wallet, error) {
	w, err := u.repository.IncreaseBalance(amount, wallet)
	if err != nil {
		return nil, err
	}

	return w, nil
}
func (u *walletUsecaseImplementation) DecreaseBalance(amount decimal.Decimal, wallet entity.Wallet) (*entity.Wallet, error) {
	w, err := u.repository.DecreaseBalance(amount, wallet)
	if err != nil {
		return nil, err
	}

	return w, nil
}
