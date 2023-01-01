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

func TestRegister(t *testing.T) {
	t.Run("Should return success", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthUsecase(usecase.AuthUsecaseImplementationConfig{
			AuthUsecase:   auth,
			UserUsecase:   userUsecase,
			WalletUsecase: walletUsecase,
		})

		req := dto.RegisterRequest{
			Fullname: "Giwang Dwi Kintan",
			Email:    "gidwikintan@gmail.com",
			Password: "password",
			Address: "Jl. Kebon Jeruk",
			CityId: 1,


		}
		res := &dto.RegisterResponse{
			Fullname: "Giwang Dwi Kintan",
			Email:    "gidwikintan@gmail.com",
			Address: "Jl. Kebon Jeruk",
			CityId: 1,
			Role: "user",
		}

		
		hashedPwd := "$2a$04$Oxfvvna0lq3qToeWXHVj.esXeXleZnTYUBAkhr55JcN73o.J33u7W"


		auth.On("HashAndSalt", req.Password).Return(hashedPwd, nil)
		userUsecase.On("CreateUser", entity.User{}).Return(entity.User{}, nil)
		user, err := uc.Register(req)

		assert.Equal(t, res, user)
		assert.Nil(t, err)
	})

	t.Run("Test case 2: error hashing password", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthUsecase(usecase.AuthUsecaseImplementationConfig{
			AuthUsecase:   auth,
			UserUsecase:   userUsecase,
			WalletUsecase: walletUsecase,
		})

		req := dto.RegisterRequest{
			Email:    "gidwikintan@gmail.com",
			Password: "password",
		}

		auth.On("HashAndSalt", req.Password).Return("", errors.New("error"))
		user, err := uc.Register(req)

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("Test case 3: error creating wallet", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthUsecase(usecase.AuthUsecaseImplementationConfig{
			AuthUsecase:   auth,
			UserUsecase:   userUsecase,
			WalletUsecase: walletUsecase,
		})

		req := dto.RegisterRequest{
			Email:    "gidwikintan@gmail.com",
			Password: "password",
		}

		hashedPwd := "$2a$04$Oxfvvna0lq3qToeWXHVj.esXeXleZnTYUBAkhr55JcN73o.J33u7W"

		auth.On("HashAndSalt", req.Password).Return(hashedPwd, nil)
		walletUsecase.On("CreateWallet").Return(nil, errors.New("error"))
		user, err := uc.Register(req)

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("Test case 4: Invalid user", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthUsecase(usecase.AuthUsecaseImplementationConfig{
			AuthUsecase:   auth,
			UserUsecase:   userUsecase,
			WalletUsecase: walletUsecase,
		})

		req := dto.RegisterRequest{
			Email:    "gidwikintan@gmail.com",
			Password: "password",
		}

		
		hashedPwd := "$2a$04$Oxfvvna0lq3qToeWXHVj.esXeXleZnTYUBAkhr55JcN73o.J33u7W"

		

		auth.On("HashAndSalt", req.Password).Return(hashedPwd, nil)
		user, err := uc.Register(req)

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

}

func TestLogin(t *testing.T) {
	t.Run("Test case 1: Valid user", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthUsecase(usecase.AuthUsecaseImplementationConfig{
			AuthUsecase:   auth,
			UserUsecase:   userUsecase,
			WalletUsecase: walletUsecase,
		})

		req := dto.LoginRequest{
			Email:    "gidwikintan@gmail.com",
			Password: "password",
		}
		res := &dto.LoginResponse{
			AccessToken: "token",
		}

		userEntity := entity.User{
			Email:    req.Email,
			Password: req.Password,
		}

		userUsecase.On("GetUserByEmail", req.Email).Return(&userEntity, nil)
		auth.On("ComparePassword", req.Password, userEntity.Password).Return(true)
		auth.On("GenerateAccessToken", &userEntity).Return("token", nil)
		user, err := uc.Login(req)

		assert.Equal(t, res.AccessToken, user)
		assert.Nil(t, err)
	})
	t.Run("Test case 2: invalid email", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthUsecase(usecase.AuthUsecaseImplementationConfig{
			AuthUsecase:   auth,
			UserUsecase:   userUsecase,
			WalletUsecase: walletUsecase,
		})

		req := dto.LoginRequest{
			Email:    "gidwikintan@gmail.com",
			Password: "password",
		}
		userUsecase.On("GetUserByEmail", req.Email).Return(nil, errors.New("error"))
		user, err := uc.Login(req)

		assert.Equal(t, "", user)
		assert.NotNil(t, err)
	})
	t.Run("Test case 3: invalid password", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthUsecase(usecase.AuthUsecaseImplementationConfig{
			AuthUsecase:   auth,
			UserUsecase:   userUsecase,
			WalletUsecase: walletUsecase,
		})

		req := dto.LoginRequest{
			Email:    "gidwikintan@gmail.com",
			Password: "password",
		}

		userEntity := entity.User{
			Email:    req.Email,
			Password: req.Password,
		}

		userUsecase.On("GetUserByEmail", req.Email).Return(&userEntity, nil)
		auth.On("ComparePassword", req.Password, userEntity.Password).Return(false)
		auth.On("GenerateAccessToken", &userEntity).Return("token", nil)
		user, err := uc.Login(req)

		assert.Equal(t, "", user)
		assert.NotNil(t, err)
	})
	t.Run("Test case 1: failed generate token", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthUsecase(usecase.AuthUsecaseImplementationConfig{
			AuthUsecase:   auth,
			UserUsecase:   userUsecase,
			WalletUsecase: walletUsecase,
		})

		req := dto.LoginRequest{
			Email:    "gidwikintan@gmail.com",
			Password: "password",
		}

		userEntity := entity.User{
			Email:    req.Email,
			Password: req.Password,
		}

		userUsecase.On("GetUserByEmail", req.Email).Return(&userEntity, nil)
		auth.On("ComparePassword", req.Password, userEntity.Password).Return(true)
		auth.On("GenerateAccessToken", &userEntity).Return("", errors.New("error"))
		user, err := uc.Login(req)

		assert.Equal(t, "", user)
		assert.NotNil(t, err)
	})

}
