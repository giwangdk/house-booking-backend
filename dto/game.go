package dto

import (
	"final-project-backend/entity"

	"github.com/shopspring/decimal"
)

type GameDetail struct {
	ID 		uint    `json:"id"`
	Chance       decimal.Decimal `json:"chance"`
	GamesPlayed  int     `json:"games_played"`
}

func (c *GameDetail) BuildResponse(game entity.Game) *GameDetail {
	return &GameDetail{
		ID: game.ID,
		Chance: game.Chance,
		GamesPlayed: game.GamesPlayed,
	}
}