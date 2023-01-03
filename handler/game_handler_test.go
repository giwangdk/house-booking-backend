package handler_test

import (
	"errors"
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/handler"
	"final-project-backend/mocks"
	"final-project-backend/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetGameByUserID(t *testing.T){
	t.Run("Should Return status ok when list of games return", func(t *testing.T) {
		uc:= new(mocks.GameUsecase)

		var games = dto.GameDetail{
			Chance: decimal.NewFromInt(1),
		}
		user := dto.UserJWT{
			ID:       1,
			Role:    "user",
		}

		h:= handler.NewHandler(handler.HandlerConfig{
			GameUsecase: uc,
		})

		uc.On("GetGameByUserID", 1).Return(&games, nil)
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/game", utils.MakeRequestBody(user))
		gin.SetMode(gin.TestMode)
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Set("user", user)
		ctx.Request = req
		h.GetGameByUserID(ctx)

		expectedRes := `{"data":{"id":0,"chance":"1","total_games_played":"0"},"status_code":200}`

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedRes, rec.Body.String())
	})
	t.Run("Should Return internal server error when fail get list of games ", func(t *testing.T) {
		uc:= new(mocks.GameUsecase)

	
		user := dto.UserJWT{
			ID:       1,
			Role:    "user",
		}

		h:= handler.NewHandler(handler.HandlerConfig{
			GameUsecase: uc,
		})

		uc.On("GetGameByUserID", 1).Return(nil, errors.New("error"))
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/game", utils.MakeRequestBody(user))
		gin.SetMode(gin.TestMode)
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Set("user", user)
		ctx.Request = req
		h.GetGameByUserID(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `{"error":"error"}`, rec.Body.String())
	})
}

func TestUpdateGame(t *testing.T){
	t.Run("Should Return status ok when list of games return", func(t *testing.T) {
		uc:= new(mocks.GameUsecase)

		
		user := dto.UserJWT{
			ID:       1,
			Role:    "user",
		}

		h:= handler.NewHandler(handler.HandlerConfig{
			GameUsecase: uc,
		})

		game:= dto.PlayGame{
			IsWin: false,
		}

		gameRes:= entity.Game{
			Chance: decimal.NewFromInt(1),
		}

		uc.On("UpdateGame", 1, game).Return(&gameRes,nil)

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/game", utils.MakeRequestBody(game))
		gin.SetMode(gin.TestMode)
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Set("user", user)
		ctx.Request = req

		h.UpdateGame(ctx)

		expectedRes := `{"data":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"chance":"1","total_games_played":"0","user_id":0},"status_code":200}`

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedRes, rec.Body.String())
	})

	t.Run("Should Return Bad request error", func(t *testing.T) {
		uc:= new(mocks.GameUsecase)

		
		user := dto.UserJWT{
			ID:       1,
			Role:    "user",
		}

		h:= handler.NewHandler(handler.HandlerConfig{
			GameUsecase: uc,
		})

		game:= dto.PlayGame{
			IsWin: false,
		}

		uc.On("UpdateGame", 1, game).Return(nil,errors.New("error"))

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/game", utils.MakeRequestBody(game))
		gin.SetMode(gin.TestMode)
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Set("user", user)
		ctx.Request = req

		h.UpdateGame(ctx)

		expectedRes := `{"status_code":500,"code":"INTERNAL_SERVER_ERROR","message":"error"}{"data":null,"status_code":200}`

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, expectedRes, rec.Body.String())
	})
	
}
