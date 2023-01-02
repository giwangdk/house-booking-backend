package repository

import (
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type GameRepository interface {
	CreateGame(userId int) (*entity.Game, error)
	IncreaseChance(chance int, Game entity.Game) (*entity.Game, error)
	DecreaseChance(chance int, Game entity.Game) (*entity.Game, error)
	GetGameByUserID(userId int) (*entity.Game, error)
	IncreaseTotalGamesPlayed(Game entity.Game) (*entity.Game, error)
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

	tx:= r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATING_Game")
	}

	u := entity.Game{
		Chance: decimal.NewFromInt(0),
		TotalGamesPlayed: decimal.NewFromInt(0),
		UserId: userId,
	}

	err := r.db.Debug().Create(&u).Error
	if err != nil {
		tx.Rollback()
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


func (r *postgresGameRepository) IncreaseChance(chance int, Game entity.Game) (*entity.Game, error) {
	tx:= r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATING_Game")
	}
	err := r.db.Model(&Game).Where("id = ?", Game.ID).Update("chance", Game.Chance.Add(decimal.NewFromInt(int64(chance)))).Error
	if err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_UPDATING_Game")
	}

	return &Game, nil
}

func (r *postgresGameRepository) DecreaseChance(chance int, Game entity.Game) (*entity.Game, error) {
	tx:= r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATING_Game")
	}

	err := r.db.Model(&Game).Where("id = ?", Game.ID).Update("chance", Game.Chance.Sub(decimal.NewFromInt(int64(chance)))).Error
	if err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_UPDATING_Game")
	}

	return &Game, nil
}

func (r *postgresGameRepository) IncreaseTotalGamesPlayed(Game entity.Game) (*entity.Game, error) {
	tx:= r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATING_Game")
	}

	err := r.db.Model(&Game).Where("id = ?", Game.ID).Update("total_games_played", Game.TotalGamesPlayed.Add(decimal.NewFromInt(1))).Error
	if err != nil {
		tx.Rollback()
		return nil, httperror.BadRequestError(err.Error(), "ERROR_UPDATING_Game")
	}

	return &Game, nil
}