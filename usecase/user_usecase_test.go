package usecase_test

import (
	"errors"
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"final-project-backend/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	t.Run("Should return success when user is created", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})
		user := entity.User{
			Email:    "",
			Role:    "admin",
		}

		repo.On("CreateUserAdmin", user).Return(&user, nil)
		res, err := uc.CreateUser(user)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user, *res)
	})

	t.Run("Should return success when user is created", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})
		user := entity.User{
			Email:    "",
			Role:    "user",
		}

		repo.On("CreateUser", user).Return(&user, nil)
		res, err := uc.CreateUser(user)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user, *res)
	})

	t.Run("Should return error when failed create admin user", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})
		user := entity.User{
			Email:    "",
			Role:    "admin",
		}

		repo.On("CreateUserAdmin", user).Return(nil, errors.New("error"))
		res, err := uc.CreateUser(user)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("Should return error when failed create  user", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})
		user := entity.User{
			Email:    "",
			Role:    "user",
		}

		repo.On("CreateUser", user).Return(nil, errors.New("error"))
		res, err := uc.CreateUser(user)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
	}

func TestUpdateRole(t *testing.T){
	t.Run("Should return success when role is updated", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})
		user := entity.User{
			Email:    "",
			Role:    "admin",
		}

		repo.On("UpdateRole", "gidwikintan@gmail.com", "host").Return(&user, nil)
		res, err := uc.UpdateRole("gidwikintan@gmail.com", "host")

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user, *res)
	})

	t.Run("Should return error when failed update role", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})


		repo.On("UpdateRole", "gidwikintan@gmail.com", "host").Return(nil, errors.New("error"))
		res, err := uc.UpdateRole("gidwikintan@gmail.com", "host")

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
}

func TestIsUserExist(t *testing.T){
	t.Run("Should return success when user is exist", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})
		user := entity.User{
			Email:    "gidwikintan@gmail.com",
			Role:    "admin",
		}

		repo.On("GetUserByEmail", "gidwikintan@gmail.com").Return(&user, nil)
		res, isExist := uc.IsUserExist("gidwikintan@gmail.com")

		assert.True(t, isExist)
		assert.NotNil(t, user)
		assert.Equal(t, user, *res)
	})
	t.Run("Should return error when user is not exist", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})

		repo.On("GetUserByEmail", "gidwikintan@gmail.com").Return(nil, errors.New("error"))
		res, isExist := uc.IsUserExist("gidwikintan@gmail.com")

		assert.False(t, isExist)
		assert.Nil(t, res)
	})
}

func TestGetUserByEmail(t *testing.T){
	t.Run("Should return success when user is exist", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})
		user := entity.User{
			Email:    "gidwikintan@gmail.com",
			Role:    "admin",
		}

		repo.On("GetUserByEmail", "gidwikintan@gmail.com").Return(&user, nil)
		res, err := uc.GetUserByEmail("gidwikintan@gmail.com")

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user, *res)
	})

	t.Run("Should return error when user is not exist", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})

		repo.On("GetUserByEmail", "gidwikintan@gmail.com").Return(nil, errors.New("error"))
		res, err := uc.GetUserByEmail("gidwikintan@gmail.com")

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
	
}

func TestGetUser(t *testing.T){
	t.Run("Should return success when user is exist", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})
		user := entity.User{
			Email:    "",
		}

		userDetail := dto.UserDetail{
			Email:   "",
		}

		repo.On("GetUser", 1).Return(&user, nil)
		res, err := uc.GetUser(1)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userDetail, *res)
	})

	t.Run("Should return error when user is not exist", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})

		repo.On("GetUser", 1).Return(nil, errors.New("error"))
		res, err := uc.GetUser(1)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
}

func TestUpdateUser(t *testing.T){
	t.Run("Should return success when user is exist", func(t *testing.T) {
		repo := new(mocks.UserRepository)
		auth := new(mocks.AuthUtil)
		uc := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
			Repository:  repo,
			AuthUsecase: auth,
		})
		user := entity.User{
			Fullname: "Gidwik Intan",
		}

		req := dto.UpdateUserRequest{
			Fullname: "Gidwik Intan",

		}

		resEx:= dto.UpdateUserResponse{
			Fullname: "Gidwik Intan",
		}

		repo.On("GetUser", 1).Return(&user, nil)
		repo.On("UpdateUser", user, 1).Return(&user, nil)
		res, err := uc.UpdateUser( req, 1)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, resEx, *res)
	})

}
