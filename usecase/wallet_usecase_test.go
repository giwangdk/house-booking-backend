package usecase_test

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"final-project-backend/usecase"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCreateWallet(t *testing.T){
	t.Run("Should return success when wallet is created", func(t *testing.T) {
		repo:= new(mocks.WalletRepository)
		uc:= usecase.NewWalletUseCase(usecase.WalletUsecaseImplementationConfig{
			Repository: repo,
		})

		wallet:= entity.Wallet{
			Balance: decimal.NewFromInt(1000),
			UserId: 1,
		}
		repo.On("CreateWallet", 1).Return(&wallet, nil)
		res, err := uc.CreateWallet(1)

		assert.Nil(t, err)
		assert.NotNil(t, wallet)
		assert.Equal(t, wallet, *res)
	})
	t.Run("Should return error when wallet is failed created", func(t *testing.T) {
		repo:= new(mocks.WalletRepository)
		uc:= usecase.NewWalletUseCase(usecase.WalletUsecaseImplementationConfig{
			Repository: repo,
		})


		repo.On("CreateWallet", 1).Return(nil, errors.New("error"))
		res, err := uc.CreateWallet(1)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Failed to create wallet", err.Error())
	})
}

func TestIsValidBalance(t *testing.T){
	t.Run("Should return true when balance is valid", func(t *testing.T) {
		repo:= new(mocks.WalletRepository)
		uc:= usecase.NewWalletUseCase(usecase.WalletUsecaseImplementationConfig{
			Repository: repo,
		})

		wallet:= entity.Wallet{
			Balance: decimal.NewFromInt(1000),
			UserId: 1,
		}

		repo.On("IsValidBalance", decimal.NewFromInt(500), wallet).Return(true, nil)
		isValid:= uc.IsValidBalance(decimal.NewFromInt(500), wallet)

		assert.True(t, isValid)
	})
}

func TestGetWalletByUserID(t *testing.T){
	t.Run("Should return success when wallet is found", func(t *testing.T) {
		repo:= new(mocks.WalletRepository)
		uc:= usecase.NewWalletUseCase(usecase.WalletUsecaseImplementationConfig{
			Repository: repo,
		})

		wallet:= entity.Wallet{
			Balance: decimal.NewFromInt(1000),
			UserId: 1,
		}

		repo.On("GetWalletByUserID", 1).Return(&wallet, nil)
		res, err := uc.GetWalletByUserID(1)

		assert.Nil(t, err)
		assert.NotNil(t, wallet)
		assert.Equal(t, wallet, *res)
	})
	t.Run("Should return error when wallet is not found", func(t *testing.T) {
		repo:= new(mocks.WalletRepository)
		uc:= usecase.NewWalletUseCase(usecase.WalletUsecaseImplementationConfig{
			Repository: repo,
		})

		repo.On("GetWalletByUserID", 1).Return(nil, errors.New("error"))
		res, err := uc.GetWalletByUserID(1)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Wallet is not found!", err.Error())
	})
}

func TestIncreaseBalance(t *testing.T){
	t.Run("Should return success when balance is increased", func(t *testing.T) {
		repo:= new(mocks.WalletRepository)
		uc:= usecase.NewWalletUseCase(usecase.WalletUsecaseImplementationConfig{
			Repository: repo,
		})

		wallet:= entity.Wallet{
			Balance: decimal.NewFromInt(1000),
			UserId: 1,
		}

		repo.On("IncreaseBalance", decimal.NewFromInt(500), wallet).Return(&wallet, nil)
		res, err := uc.IncreaseBalance(decimal.NewFromInt(500), wallet)

		assert.Nil(t, err)
		assert.NotNil(t, wallet)
		assert.Equal(t, wallet, *res)
	})
	t.Run("Should return error when balance is failed to increase", func(t *testing.T) {
		repo:= new(mocks.WalletRepository)
		uc:= usecase.NewWalletUseCase(usecase.WalletUsecaseImplementationConfig{
			Repository: repo,
		})

		wallet:= entity.Wallet{
			Balance: decimal.NewFromInt(1000),
			UserId: 1,
		}

		repo.On("IncreaseBalance", decimal.NewFromInt(500), wallet).Return(nil, errors.New("error"))
		res, err := uc.IncreaseBalance(decimal.NewFromInt(500), wallet)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Failed to increase balance", err.Error())
	})
}

func TestDecreaseBalance(t *testing.T){
	t.Run("Should return success when balance is decreased", func(t *testing.T) {
		repo:= new(mocks.WalletRepository)
		uc:= usecase.NewWalletUseCase(usecase.WalletUsecaseImplementationConfig{
			Repository: repo,
		})

		wallet:= entity.Wallet{
			Balance: decimal.NewFromInt(1000),
			UserId: 1,
		}

		repo.On("DecreaseBalance", decimal.NewFromInt(500), wallet).Return(&wallet, nil)
		res, err := uc.DecreaseBalance(decimal.NewFromInt(500), wallet)

		assert.Nil(t, err)
		assert.NotNil(t, wallet)
		assert.Equal(t, wallet, *res)
	})
	t.Run("Should return error when balance is failed to decrease", func(t *testing.T) {
		repo:= new(mocks.WalletRepository)
		uc:= usecase.NewWalletUseCase(usecase.WalletUsecaseImplementationConfig{
			Repository: repo,
		})

		wallet:= entity.Wallet{
			Balance: decimal.NewFromInt(1000),
			UserId: 1,
		}

		repo.On("DecreaseBalance", decimal.NewFromInt(500), wallet).Return(nil, errors.New("error"))
		res, err := uc.DecreaseBalance(decimal.NewFromInt(500), wallet)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Failed to decrease balance", err.Error())
	})
}