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


func TestCreateGame(t *testing.T) {
	t.Run("Should return success when games created", func(t *testing.T) {

		gameRepo := new(mocks.GameRepository)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
			Repository: gameRepo,
			WalletUsecase: walletUsecase,
			WalletTxRepo: walletTxRepo,
		})
		games:= entity.Game{
			UserId: 1,
			Chance: decimal.NewFromInt(0),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}
		gameRepo.On("CreateGame", 1).Return(&games, nil)
		res, err := uc.CreateGame(1)

		assert.Nil(t, err)
		assert.Equal(t, games, *res)	
	})

	t.Run("Should return error when games created", func(t *testing.T) {
		
		gameRepo := new(mocks.GameRepository)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
			Repository: gameRepo,
			WalletUsecase: walletUsecase,
			WalletTxRepo: walletTxRepo,
		})
		gameRepo.On("CreateGame", 1).Return(nil, errors.New("error"))
		res, err := uc.CreateGame(1)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
}

func TestUpdateGame(t *testing.T) {
	t.Run("Should return success when games updated", func(t *testing.T) {

		gameRepo := new(mocks.GameRepository)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
			Repository: gameRepo,
			WalletUsecase: walletUsecase,
			WalletTxRepo: walletTxRepo,
		})

		req := dto.PlayGame{
			IsWin: true,
		}
		games:= entity.Game{
			UserId: 1,
			Chance: decimal.NewFromInt(10),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}

		wallet:= entity.Wallet{
			UserId: 1,
			Balance: decimal.NewFromInt(0),
		}

		walletTx:= entity.WalletTransaction{
			Sender: 0,
			Recipient: 0,
			Amount: decimal.NewFromInt(100000),
			Description: "Redeem Money from game",
		}

		gameRepo.On("GetGameByUserID", 1).Return(&games, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&wallet, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(100000), wallet).Return(&wallet, nil)
		walletTxRepo.On("CreateWalletTransaction", walletTx).Return(&walletTx, nil)
		gameRepo.On("DecreaseChance", 1, games).Return(&games, nil)
		gameRepo.On("IncreaseTotalGamesPlayed",  games).Return(&games, nil)
		game, err := uc.UpdateGame(1,req )

		assert.Nil(t, err)
		assert.Equal(t, games, *game)	
	})


	t.Run("Should return error when game not found", func(t *testing.T) {

		gameRepo := new(mocks.GameRepository)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
			Repository: gameRepo,
			WalletUsecase: walletUsecase,
			WalletTxRepo: walletTxRepo,
		})

		req := dto.PlayGame{
			IsWin: true,
		}
		games:= entity.Game{
			UserId: 1,
			Chance: decimal.NewFromInt(10),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}

		gameRepo.On("GetGameByUserID", 1).Return(&games, errors.New("error"))
		game, err := uc.UpdateGame(1,req )

		assert.NotNil(t, err)
		assert.Nil(t, game)
		assert.Equal(t, "error", err.Error())
	})
	


	t.Run("Should return error increase balance fail when games updated", func(t *testing.T) {

		gameRepo := new(mocks.GameRepository)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
			Repository: gameRepo,
			WalletUsecase: walletUsecase,
			WalletTxRepo: walletTxRepo,
		})

		req := dto.PlayGame{
			IsWin: true,
		}
		games:= entity.Game{
			UserId: 1,
			Chance: decimal.NewFromInt(10),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}

		wallet:= entity.Wallet{
			UserId: 1,
			Balance: decimal.NewFromInt(0),
		}


		gameRepo.On("GetGameByUserID", 1).Return(&games, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&wallet, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(100000), wallet).Return(nil, errors.New("error"))
		game, err := uc.UpdateGame(1,req )

		assert.NotNil(t, err)
		assert.Nil(t, game)
		assert.Equal(t, "error", err.Error())
	})

	t.Run("Should return success when games updated", func(t *testing.T) {

		gameRepo := new(mocks.GameRepository)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
			Repository: gameRepo,
			WalletUsecase: walletUsecase,
			WalletTxRepo: walletTxRepo,
		})

		req := dto.PlayGame{
			IsWin: true,
		}
		games:= entity.Game{
			UserId: 1,
			Chance: decimal.NewFromInt(10),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}

		wallet:= entity.Wallet{
			UserId: 1,
			Balance: decimal.NewFromInt(0),
		}

		walletTx:= entity.WalletTransaction{
			Sender: 0,
			Recipient: 0,
			Amount: decimal.NewFromInt(100000),
			Description: "Redeem Money from game",
		}

		gameRepo.On("GetGameByUserID", 1).Return(&games, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&wallet, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(100000), wallet).Return(&wallet, nil)
		walletTxRepo.On("CreateWalletTransaction", walletTx).Return(&walletTx, nil)
		gameRepo.On("DecreaseChance", 1, games).Return(&games, nil)
		gameRepo.On("IncreaseTotalGamesPlayed",  games).Return(&games, nil)
		game, err := uc.UpdateGame(1,req )

		assert.Nil(t, err)
		assert.Equal(t, games, *game)	
	})


}

func TestGetGameByUserID(t *testing.T) {
	t.Run("Should return success when game is found", func(t *testing.T) {

		gameRepo := new(mocks.GameRepository)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
			Repository: gameRepo,
			WalletUsecase: walletUsecase,
			WalletTxRepo: walletTxRepo,
		})
		games:= entity.Game{
			UserId: 1,
			Chance: decimal.NewFromInt(0),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}

		res:= dto.GameDetail{
			Chance: decimal.NewFromInt(0),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}
		gameRepo.On("GetGameByUserID", 1).Return(&games, nil)
   		game, err := uc.GetGameByUserID(1)	

		assert.Nil(t, err)
		assert.Equal(t, &res, game)
	})

	t.Run("Should return error when game is not found", func(t *testing.T) {
		
		gameRepo := new(mocks.GameRepository)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
			Repository: gameRepo,
			WalletUsecase: walletUsecase,
			WalletTxRepo: walletTxRepo,
		})
		games:= entity.Game{
			UserId: 1,
			Chance: decimal.NewFromInt(0),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}

	
		gameRepo.On("GetGameByUserID", 1).Return(&games, errors.New("error"))
   		game, err := uc.GetGameByUserID(1)	

		assert.NotNil(t, err)
		assert.Nil(t, game)
		assert.Equal(t, "error", err.Error())
	})

}


func TestIncreaseChance(t *testing.T) {
	t.Run("Should return success when chacne is increased", func(t *testing.T) {

		gameRepo := new(mocks.GameRepository)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
			Repository: gameRepo,
			WalletUsecase: walletUsecase,
			WalletTxRepo: walletTxRepo,
		})
		games:= entity.Game{
			UserId: 1,
			Chance: decimal.NewFromInt(0),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}


		gameRepo.On("IncreaseChance", 1, games).Return(&games, nil)
		game, err := uc.IncreaseChance(1, games)

		assert.Nil(t, err)
		assert.Equal(t, &games, game)
	})

	t.Run("Should return error when chance is not increased", func(t *testing.T) {
		
		gameRepo := new(mocks.GameRepository)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
			Repository: gameRepo,
			WalletUsecase: walletUsecase,
			WalletTxRepo: walletTxRepo,
		})
		games:= entity.Game{
			UserId: 1,
			Chance: decimal.NewFromInt(0),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}

		gameRepo.On("IncreaseChance", 1, games).Return(&games, errors.New("error"))
		game, err := uc.IncreaseChance(1, games)

		assert.NotNil(t, err)
		assert.Nil(t, game)
		assert.Equal(t, "error", err.Error())
	})
}

func TestDecreaseChance(t *testing.T) {
	t.Run("Should return success when chacne is increased", func(t *testing.T) {

		gameRepo := new(mocks.GameRepository)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
			Repository: gameRepo,
			WalletUsecase: walletUsecase,
			WalletTxRepo: walletTxRepo,
		})
		games:= entity.Game{
			UserId: 1,
			Chance: decimal.NewFromInt(0),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}


		gameRepo.On("DecreaseChance", 1, games).Return(&games, nil)
		game, err := uc.DecreaseChance(1, games)

		assert.Nil(t, err)
		assert.Equal(t, &games, game)
	})

	t.Run("Should return error when chance is not increased", func(t *testing.T) {
		
		gameRepo := new(mocks.GameRepository)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
			Repository: gameRepo,
			WalletUsecase: walletUsecase,
			WalletTxRepo: walletTxRepo,
		})
		games:= entity.Game{
			UserId: 1,
			Chance: decimal.NewFromInt(0),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}

		gameRepo.On("DecreaseChance", 1, games).Return(&games, errors.New("error"))
		game, err := uc.DecreaseChance(1, games)

		assert.NotNil(t, err)
		assert.Nil(t, game)
		assert.Equal(t, "error", err.Error())
	})
}




	