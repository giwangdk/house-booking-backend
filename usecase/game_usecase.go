package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/httperror"
	"final-project-backend/repository"

	"github.com/shopspring/decimal"
)

type GameUsecase interface {
	CreateGame(userID int) (*entity.Game, error)
	IncreaseChance(chance int, Game entity.Game) (*entity.Game, error)
	DecreaseChance(chance int, Game entity.Game) (*entity.Game, error)
	GetGameByUserID(userId int) (*dto.GameDetail, error) 
	UpdateGame(userId int, req dto.PlayGame ) (*entity.Game, error)
}

type GameUsecaseImplementation struct {
	repository repository.GameRepository
	walletUsecase WalletUsecase
	WalletTxRepo repository.WalletTransactionRepository
}

type GameUsecaseImplementationConfig struct {
	Repository repository.GameRepository
	WalletUsecase WalletUsecase
	WalletTxRepo repository.WalletTransactionRepository
}

func NewGameUseCase(c GameUsecaseImplementationConfig) GameUsecase {
	return &GameUsecaseImplementation{
		repository: c.Repository,
		walletUsecase: c.WalletUsecase,
		WalletTxRepo: c.WalletTxRepo,
	}
}


func (u *GameUsecaseImplementation) CreateGame(userId int) (*entity.Game, error) {

	w, err := u.repository.CreateGame(userId)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (u *GameUsecaseImplementation) UpdateGame(userId int, req dto.PlayGame ) (*entity.Game, error) {
	game, err := u.repository.GetGameByUserID(userId)
	if err != nil {
		return nil, err
	}

	if game.Chance.LessThanOrEqual(decimal.NewFromInt(0)) {
		return nil, httperror.BadRequestError("You don't have enough chance to play", "FAILED_UPDATE_GAME")
	}
	
	wallet, err:= u.walletUsecase.GetWalletByUserID(userId)
	if err != nil {
		return nil, err
	}
	
	if req.IsWin{
		_, err = u.walletUsecase.IncreaseBalance(decimal.NewFromInt(100000),*wallet)
		if err != nil {
			return nil, err
		}
		entity := entity.WalletTransaction{
			Sender:      int64(wallet.ID),
			Recipient:  int64(wallet.ID),
			Amount:      decimal.NewFromInt(100000),
			Description: "Redeem Money from game",
		}
	
		
		_, err := u.WalletTxRepo.CreateWalletTransaction(entity)
		if err != nil {
			return nil, err
		}
		}else{
		_, err = u.walletUsecase.IncreaseBalance(decimal.NewFromInt(1000),*wallet)
		if err != nil {
			return nil, err
		}
		entity := entity.WalletTransaction{
			Sender:      int64(wallet.ID),
			Recipient:  int64(wallet.ID),
			Amount:      decimal.NewFromInt(1000),
			Description: "Redeem Money from game ",
		}
	
		
		_, err := u.WalletTxRepo.CreateWalletTransaction(entity)
		if err != nil {
			return nil, err
		}
	}

	if game.TotalGamesPlayed.Mod(decimal.NewFromInt(10)).Equal(decimal.NewFromInt(0)) && game.TotalGamesPlayed.GreaterThan(decimal.NewFromInt(10)) {
		game, err = u.IncreaseChance(1,*game)
		if err != nil {
			return nil, err
		}
	}


	game, err = u.DecreaseChance(1,*game)
	if err != nil {
		return nil, err
	}
	game, err = u.repository.IncreaseTotalGamesPlayed(*game)
	if err != nil {
		return nil, err
	}

	return game, nil
}



func (u *GameUsecaseImplementation) GetGameByUserID(userId int) (*dto.GameDetail, error) {
	g, err := u.repository.GetGameByUserID(userId)
	if err != nil {
		return nil, err
	}


	res:= (&dto.GameDetail{}).BuildResponse(*g)

	return res, nil
}

func (u *GameUsecaseImplementation) IncreaseChance(chance int, Game entity.Game) (*entity.Game, error) {
	w, err := u.repository.IncreaseChance(chance, Game)
	if err != nil {
		return nil, err
	}

	return w, nil
}
func (u *GameUsecaseImplementation) DecreaseChance(chance int, Game entity.Game) (*entity.Game, error) {
	w, err := u.repository.DecreaseChance(chance, Game)
	if err != nil {
		return nil, err
	}

	return w, nil
}
