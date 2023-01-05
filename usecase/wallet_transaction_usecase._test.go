package usecase_test

import (
	"errors"
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"final-project-backend/usecase"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetWalletTransactionUser(t *testing.T) {
	t.Run("Should return success when get wallet transaction user", func(t *testing.T) {
		repo := new(mocks.WalletTransactionRepository)
		walletUsecase := new(mocks.WalletUsecase)
		gameRepo := new(mocks.GameRepository)

		uc := usecase.NewWalletTransactionUseCase(usecase.WalletTransactionUsecaseImplementationConfig{
			Repository:     repo,
			WalletUsecase:  walletUsecase,
			GameRepository: gameRepo,
		})

		var walletTransactionUser *[]entity.WalletTransaction
		repo.On("GetWalletTransactionsUser", 1).Return(walletTransactionUser, nil)
		res, err := uc.GetWalletTransactionsUser(1)

		assert.Nil(t, err)
		assert.Equal(t, walletTransactionUser, res)
	})
	t.Run("Should return error when get wallet transaction user", func(t *testing.T) {
		repo := new(mocks.WalletTransactionRepository)
		walletUsecase := new(mocks.WalletUsecase)
		gameRepo := new(mocks.GameRepository)

		uc := usecase.NewWalletTransactionUseCase(usecase.WalletTransactionUsecaseImplementationConfig{
			Repository:     repo,
			WalletUsecase:  walletUsecase,
			GameRepository: gameRepo,
		})

		repo.On("GetWalletTransactionsUser", 1).Return(nil, errors.New("error"))
		res, err := uc.GetWalletTransactionsUser(1)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Wallet Transaction is not found!", err.Error())
	})
}

func TestTopup(t *testing.T) {
	t.Run("Should return success when topup", func(t *testing.T) {
		repo := new(mocks.WalletTransactionRepository)
		walletUsecase := new(mocks.WalletUsecase)
		gameRepo := new(mocks.GameRepository)

		uc := usecase.NewWalletTransactionUseCase(usecase.WalletTransactionUsecaseImplementationConfig{
			Repository:     repo,
			WalletUsecase:  walletUsecase,
			GameRepository: gameRepo,
		})

		req := dto.TopUpRequest{
			Sender:    1,
			Recipient: 1,
			Amount:    decimal.NewFromInt(100000),
		}
		game := entity.Game{
			Chance: decimal.NewFromInt(1),
			UserId: 1,
		}

		wallet := entity.Wallet{
			UserId:  1,
			Balance: decimal.NewFromInt(100000),
		}

		entity := entity.WalletTransaction{
			Sender:      int64(wallet.ID),
			Recipient:   int64(wallet.ID),
			Amount:      decimal.NewFromInt(100000),
			Description: "Top Up ",
		}

		response := dto.TopUpResponse{
			Amount:      decimal.NewFromInt(100000),
			Description: "Top Up ",
		}

		gameRepo.On("GetGameByUserID", 1).Return(&game, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&wallet, nil)
		repo.On("CreateWalletTransaction", entity).Return(&entity, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(100000), wallet).Return(&wallet, nil)
		res, err := uc.TopUp(req)

		assert.Nil(t, err)
		assert.Equal(t, &response, res)
	})
	t.Run("Should return error  when failed topup because game is not found", func(t *testing.T) {
		repo := new(mocks.WalletTransactionRepository)
		walletUsecase := new(mocks.WalletUsecase)
		gameRepo := new(mocks.GameRepository)

		uc := usecase.NewWalletTransactionUseCase(usecase.WalletTransactionUsecaseImplementationConfig{
			Repository:     repo,
			WalletUsecase:  walletUsecase,
			GameRepository: gameRepo,
		})

		req := dto.TopUpRequest{
			Sender:    1,
			Recipient: 1,
			Amount:    decimal.NewFromInt(100000),
		}

		gameRepo.On("GetGameByUserID", 1).Return(nil, errors.New("error"))

		res, err := uc.TopUp(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})

	t.Run("Should return error when wallet not found", func(t *testing.T) {
		repo := new(mocks.WalletTransactionRepository)
		walletUsecase := new(mocks.WalletUsecase)
		gameRepo := new(mocks.GameRepository)

		uc := usecase.NewWalletTransactionUseCase(usecase.WalletTransactionUsecaseImplementationConfig{
			Repository:     repo,
			WalletUsecase:  walletUsecase,
			GameRepository: gameRepo,
		})

		req := dto.TopUpRequest{
			Sender:    1,
			Recipient: 1,
			Amount:    decimal.NewFromInt(100000),
		}
		game := entity.Game{
			Chance: decimal.NewFromInt(1),
			UserId: 1,
		}
		gameRepo.On("GetGameByUserID", 1).Return(&game, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(nil, errors.New("error"))

		res, err := uc.TopUp(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("Should return error when fail create wallet transaction", func(t *testing.T) {
		repo := new(mocks.WalletTransactionRepository)
		walletUsecase := new(mocks.WalletUsecase)
		gameRepo := new(mocks.GameRepository)

		uc := usecase.NewWalletTransactionUseCase(usecase.WalletTransactionUsecaseImplementationConfig{
			Repository:     repo,
			WalletUsecase:  walletUsecase,
			GameRepository: gameRepo,
		})

		req := dto.TopUpRequest{
			Sender:    1,
			Recipient: 1,
			Amount:    decimal.NewFromInt(100000),
		}
		game := entity.Game{
			Chance: decimal.NewFromInt(1),
			UserId: 1,
		}

		wallet := entity.Wallet{
			UserId:  1,
			Balance: decimal.NewFromInt(100000),
		}

		entity := entity.WalletTransaction{
			Sender:      int64(wallet.ID),
			Recipient:   int64(wallet.ID),
			Amount:      decimal.NewFromInt(100000),
			Description: "Top Up ",
		}

		gameRepo.On("GetGameByUserID", 1).Return(&game, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&wallet, nil)
		repo.On("CreateWalletTransaction", entity).Return(nil, errors.New("error"))
		res, err := uc.TopUp(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Error creating wallet transaction", err.Error())
	})
	t.Run("Should return error when failed increase balance", func(t *testing.T) {
		repo := new(mocks.WalletTransactionRepository)
		walletUsecase := new(mocks.WalletUsecase)
		gameRepo := new(mocks.GameRepository)

		uc := usecase.NewWalletTransactionUseCase(usecase.WalletTransactionUsecaseImplementationConfig{
			Repository:     repo,
			WalletUsecase:  walletUsecase,
			GameRepository: gameRepo,
		})

		req := dto.TopUpRequest{
			Sender:    1,
			Recipient: 1,
			Amount:    decimal.NewFromInt(100000),
		}
		game := entity.Game{
			Chance: decimal.NewFromInt(1),
			UserId: 1,
		}

		wallet := entity.Wallet{
			UserId:  1,
			Balance: decimal.NewFromInt(100000),
		}

		entity := entity.WalletTransaction{
			Sender:      int64(wallet.ID),
			Recipient:   int64(wallet.ID),
			Amount:      decimal.NewFromInt(100000),
			Description: "Top Up ",
		}

		gameRepo.On("GetGameByUserID", 1).Return(&game, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&wallet, nil)
		repo.On("CreateWalletTransaction", entity).Return(&entity, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(100000), wallet).Return(nil, errors.New("error"))
		res, err := uc.TopUp(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("Should return success when increase chance  if topup more than 500000", func(t *testing.T) {
		repo := new(mocks.WalletTransactionRepository)
		walletUsecase := new(mocks.WalletUsecase)
		gameRepo := new(mocks.GameRepository)

		uc := usecase.NewWalletTransactionUseCase(usecase.WalletTransactionUsecaseImplementationConfig{
			Repository:     repo,
			WalletUsecase:  walletUsecase,
			GameRepository: gameRepo,
		})

		req := dto.TopUpRequest{
			Sender:    1,
			Recipient: 1,
			Amount:    decimal.NewFromInt(500000),
		}
		game := entity.Game{
			Chance: decimal.NewFromInt(1),
			UserId: 1,
		}

		wallet := entity.Wallet{
			UserId:  1,
			Balance: decimal.NewFromInt(500000),
		}

		entity := entity.WalletTransaction{
			Sender:      int64(wallet.ID),
			Recipient:   int64(wallet.ID),
			Amount:      decimal.NewFromInt(500000),
			Description: "Top Up ",
		}

		response := dto.TopUpResponse{
			Amount:      decimal.NewFromInt(500000),
			Description: "Top Up ",
		}

		gameRepo.On("GetGameByUserID", 1).Return(&game, nil)
		gameRepo.On("IncreaseChance", 1, game).Return(&game, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&wallet, nil)
		repo.On("CreateWalletTransaction", entity).Return(&entity, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(500000), wallet).Return(&wallet, nil)
		res, err := uc.TopUp(req)

		assert.Nil(t, err)
		assert.Equal(t, decimal.NewFromInt(1), game.Chance)
		assert.Equal(t, &response, res)
	})
	t.Run("Should return error if failed increase chance  when topup more than 500000", func(t *testing.T) {
		repo := new(mocks.WalletTransactionRepository)
		walletUsecase := new(mocks.WalletUsecase)
		gameRepo := new(mocks.GameRepository)

		uc := usecase.NewWalletTransactionUseCase(usecase.WalletTransactionUsecaseImplementationConfig{
			Repository:     repo,
			WalletUsecase:  walletUsecase,
			GameRepository: gameRepo,
		})

		req := dto.TopUpRequest{
			Sender:    1,
			Recipient: 1,
			Amount:    decimal.NewFromInt(500000),
		}
		game := entity.Game{
			Chance: decimal.NewFromInt(1),
			UserId: 1,
		}

		gameRepo.On("GetGameByUserID", 1).Return(&game, nil)
		gameRepo.On("IncreaseChance", 1, game).Return(nil, errors.New("error"))
		res, err := uc.TopUp(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})

}
