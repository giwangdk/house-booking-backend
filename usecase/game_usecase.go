package usecase

import (
	"final-project-backend/entity"
	"final-project-backend/repository"

	"github.com/shopspring/decimal"
)

type GameUsecase interface {
	CreateGame(userID int) (*entity.Game, error)
	IncreaseChance(amount decimal.Decimal, Game entity.Game) (*entity.Game, error)
	DecreaseChance(amount decimal.Decimal, Game entity.Game) (*entity.Game, error)
	GetGameByID(id int) (*entity.Game, error)
}

type GameUsecaseImplementation struct {
	repository repository.GameRepository
}

type GameUsecaseImplementationConfig struct {
	Repository repository.GameRepository
}

func NewGameUseCase(c GameUsecaseImplementationConfig) GameUsecase {
	return &GameUsecaseImplementation{
		repository: c.Repository,
	}
}

func (u *GameUsecaseImplementation) CreateGame(userId int) (*entity.Game, error) {

	w, err := u.repository.CreateGame(userId)
	if err != nil {
		return nil, err
	}

	return w, nil
}


func (u *GameUsecaseImplementation) GetGameByID(id int) (*entity.Game, error) {
	w, err := u.repository.GetGameByID(id)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (u *GameUsecaseImplementation) IncreaseChance(amount decimal.Decimal, Game entity.Game) (*entity.Game, error) {
	w, err := u.repository.IncreaseChance(amount, Game)
	if err != nil {
		return nil, err
	}

	return w, nil
}
func (u *GameUsecaseImplementation) DecreaseChance(amount decimal.Decimal, Game entity.Game) (*entity.Game, error) {
	w, err := u.repository.DecreaseChance(amount, Game)
	if err != nil {
		return nil, err
	}

	return w, nil
}
