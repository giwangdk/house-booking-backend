package handler_test

import (
	"bytes"
	"encoding/json"
	"final-project-backend/dto"
	"final-project-backend/mocks"
	"final-project-backend/server"
	"final-project-backend/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterAdmin(t *testing.T) {
	t.Run("Should Return Bad Request Error when the email field is null", func(t *testing.T) {
		uc := new(mocks.AuthAdminUsecase)
		user := dto.RegisterRequest{
			Fullname:     "giwang",
			Password: "password",
			CityId:1,
		}
		uc.On("Register", user).Return(nil, http.StatusBadRequest)
		handler := &server.RouterConfig{
			AuthAdminUsecase: uc,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)
		expectedRes := `{"status_code":400,"code":"ERROR_REGISTER","message":"Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`

		req, _ := http.NewRequest("POST", "/admin/register", &buf)
		_, res := utils.ServeReq(handler, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expectedRes, res.Body.String())
	})
	t.Run("Should Return Bad Request Error when the city_id field is null", func(t *testing.T) {
		uc := new(mocks.AuthAdminUsecase)
		user := dto.RegisterRequest{
			Fullname:     "giwang",
			Password: "password",
			Email:    "gidiwkintan@gmail.com",
		}
		uc.On("Register", user).Return(nil, http.StatusBadRequest)
		handler := &server.RouterConfig{
			AuthAdminUsecase: uc,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)
		expectedRes := `{"status_code":400,"code":"ERROR_REGISTER","message":"Key: 'RegisterRequest.CityId' Error:Field validation for 'CityId' failed on the 'required' tag"}`

		req, _ := http.NewRequest("POST", "/admin/register", &buf)
		_, res := utils.ServeReq(handler, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expectedRes, res.Body.String())
	})
	t.Run("Should Return Bad Request Error when the fullname field is null", func(t *testing.T) {
		uc := new(mocks.AuthAdminUsecase)
		user := dto.RegisterRequest{
			Password: "password",
			Email:    "gidiwkintan@gmail.com",
			CityId:1,
		}
		uc.On("Register", user).Return(nil, http.StatusBadRequest)
		handler := &server.RouterConfig{
			AuthAdminUsecase: uc,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)
		expectedRes := `{"status_code":400,"code":"ERROR_REGISTER","message":"Key: 'RegisterRequest.Fullname' Error:Field validation for 'Fullname' failed on the 'required' tag"}`

		req, _ := http.NewRequest("POST", "/admin/register", &buf)
		_, res := utils.ServeReq(handler, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expectedRes, res.Body.String())
	})
	t.Run("Should Return Internal Server Error when fail register user", func(t *testing.T) {
		uc := new(mocks.AuthAdminUsecase)
		user := dto.RegisterRequest{
			Fullname:     "giwang",
			Email:    "gidiwkintan@gmail.com",
			Password: "password",
			CityId:1,
		}
		uc.On("Register", user).Return(nil, http.StatusInternalServerError)
		handler := &server.RouterConfig{
			AuthAdminUsecase: uc,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)

		req, _ := http.NewRequest("POST", "/admin/register", &buf)
		_, res := utils.ServeReq(handler, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "", res.Body.String())
	})
	t.Run("Should Return Status Success when user is created", func(t *testing.T) {
		uc := new(mocks.AuthAdminUsecase)
		user := dto.RegisterRequest{
			Fullname:     "giwang",
			Email:    "gidiwkintan@gmail.com",
			Password: "password",
			CityId:1,
		}

		userRes := dto.RegisterResponse{
			Fullname:     user.Fullname,
			Email:    user.Email,
			CityId: 1,
			Role: "admin",

		}

		expectedRes := `{"data":{"fullname":"giwang","email":"gidiwkintan@gmail.com","address":"","city_id":1,"role":"admin"},"status_code":201}`

		uc.On("Register", user).Return(&userRes, nil)
		handler := &server.RouterConfig{
			AuthAdminUsecase: uc,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)

		req, _ := http.NewRequest("POST", "/admin/register", &buf)
		_, res := utils.ServeReq(handler, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, expectedRes, res.Body.String())
	})
}

func TestLoginAdmin(t *testing.T) {
	t.Run("Should Return Bad Request Error when the email field is null", func(t *testing.T) {
		uc := new(mocks.AuthAdminUsecase)
		user := dto.LoginRequest{
			Password: "password",
		}
		uc.On("Register", user).Return(nil, http.StatusBadRequest)
		handler := &server.RouterConfig{
			AuthAdminUsecase: uc,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)
		expectedRes := `{"status_code":400,"code":"ERROR_LOGIN","message":"Key: 'LoginRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`

		req, _ := http.NewRequest("POST", "/admin/login", &buf)
		_, res := utils.ServeReq(handler, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expectedRes, res.Body.String())
	})
	t.Run("Should Return Internal server Error when the failed login", func(t *testing.T) {
		uc := new(mocks.AuthAdminUsecase)

		user := dto.LoginRequest{
			Email:    "gidwikintan@gmail.com",
			Password: "password",
		}
		uc.On("Login", user).Return(nil, http.StatusInternalServerError)
		handler := &server.RouterConfig{
			AuthAdminUsecase: uc,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)
		

		req, _ := http.NewRequest("POST", "/admin/login", &buf)
		_, res := utils.ServeReq(handler, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "", res.Body.String())
	})
	t.Run("Should Return access token when successfully login", func(t *testing.T) {
		uc := new(mocks.AuthAdminUsecase)

		user := dto.LoginRequest{
			Email:    "gidwikintan@gmail.com",
			Password: "password",
		}

		response:= dto.LoginResponse{
			AccessToken: "token",
		}

		uc.On("Login", user).Return(&response, nil)
		handler := &server.RouterConfig{
			AuthAdminUsecase: uc,
		}
		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(user)

		req, _ := http.NewRequest("POST", "/admin/login", &buf)
		_, res := utils.ServeReq(handler, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, `{"data":{"access_token":"token"},"status_code":200}`, res.Body.String())
	})
}
