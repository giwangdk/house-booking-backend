package dto

import (
	"final-project-backend/entity"

	"github.com/shopspring/decimal"
)

type GameDetail struct {
	ID 		uint    `json:"id"`
	Chance       decimal.Decimal `json:"chance"`
	TotalGamesPlayed  int     `json:"total_games_played"`
}

func (c *GameDetail) BuildResponse(game entity.Game) *GameDetail {
	return &GameDetail{
		ID: game.ID,
		Chance: game.Chance,
		TotalGamesPlayed: game.TotalGamesPlayed,
	}
}