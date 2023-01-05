package handler_test

import (
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"final-project-backend/server"
	"final-project-backend/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCities(t *testing.T){
	t.Run("Should Return status ok when list of cities return", func(t *testing.T) {
		uc:= new(mocks.CityUsecase)

		var cities []entity.City

		uc.On("GetCities").Return(&cities, nil)

		handler := &server.RouterConfig{
			CityUsecase: uc,
		}

		expectedRes := `{"data":null,"status_code":200}`
		req, _ := http.NewRequest("GET", "/cities", nil)
		_, res := utils.ServeReq(handler, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expectedRes, res.Body.String())
	})
	t.Run("Should Return internal server error when fail get list of cities ", func(t *testing.T) {
		uc:= new(mocks.CityUsecase)


		uc.On("GetCities").Return(nil, http.StatusInternalServerError)

		handler := &server.RouterConfig{
			CityUsecase: uc,
		}

		req, _ := http.NewRequest("GET", "/cities", nil)
		_, res := utils.ServeReq(handler, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "", res.Body.String())
	})
}