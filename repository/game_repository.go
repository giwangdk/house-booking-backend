package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type GameRepository interface {
	CreateGame(userId int) (*entity.Game, error)
	IncreaseChance(amount decimal.Decimal, Game entity.Game) (*entity.Game, error)
	DecreaseChance(amount decimal.Decimal, Game entity.Game) (*entity.Game, error)
	IsValidChance(amount decimal.Decimal, Game entity.Game) bool
	GetGameByUserID(userId int) (*entity.Game, error)
}

type postgresGameRepository struct {
	db *gorm.DB
}
type PostgresGameRepositoryConfig struct {
	DB *gorm.DB
}

func NewPostgresGameRepository(c PostgresGameRepositoryConfig) GameRepository {
	return &postgresGameRepository{
		db: c.DB,
	}
}

func (r *postgresGameRepository) CreateGame(userId int) (*entity.Game, error) {
	u := entity.Game{
		Chance: decimal.NewFromInt(0),
		GamesPlayed: 0,
		UserId: userId,
	}

	err := r.db.Debug().Create(&u).Error
	if err != nil {
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATING_Game")
	}

	return &u, nil
}

func (r *postgresGameRepository) GetGameByUserID(userId int) (*entity.Game, error) {
	var u entity.Game


	err := r.db.Debug().Where("user_id = ?", userId).Take(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *postgresGameRepository) IsValidChance(amount decimal.Decimal, Game entity.Game) bool {
	return Game.Chance.GreaterThanOrEqual(amount)
}

func (r *postgresGameRepository) IncreaseChance(amount decimal.Decimal, Game entity.Game) (*entity.Game, error) {
	Game.Chance = Game.Chance.Add(amount)

	err := r.db.Save(&Game).Error
	if err != nil {
		return nil, httperror.BadRequestError(err.Error(), "ERROR_UPDATING_Game")
	}

	return &Game, nil
}

func (r *postgresGameRepository) DecreaseChance(amount decimal.Decimal, Game entity.Game) (*entity.Game, error) {

	Game.Chance = Game.Chance.Sub(amount)

	err := r.db.Save(&Game).Error
	if err != nil {
		return nil, httperror.BadRequestError(err.Error(), "ERROR_UPDATING_Game")
	}

	return &Game, nil
}
