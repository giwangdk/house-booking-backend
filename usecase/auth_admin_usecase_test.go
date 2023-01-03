package usecase_test

import (
	"errors"
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"final-project-backend/usecase"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAdmin(t *testing.T) {
	t.Run("Should return success when user is successfully registered", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		gameUsecase := new(mocks.GameUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthAdminUsecase(usecase.AuthAdminUsecaseImplementationConfig{
			AuthAdminUsecase: auth,
			UserUsecase:      userUsecase,
			GameUsecase:      gameUsecase,
			WalletUsecase:    walletUsecase,
		})

		req := dto.RegisterRequest{
			Fullname: "Giwang Dwi Kintan",
			Email:    "gidwikintan@gmail.com",
			Password: "password",
			Address:  "Jl. Kebon Jeruk",
			CityId:   1,
		}
		res := &dto.RegisterResponse{
			Fullname: "Giwang Dwi Kintan",
			Email:    "gidwikintan@gmail.com",
			Address:  "Jl. Kebon Jeruk",
			CityId:   1,
			Role:     "admin",
		}

		wallet := entity.Wallet{
			UserId:  1,
			Balance: decimal.NewFromInt(0),
		}

		game := entity.Game{
			UserId:           1,
			Chance:           decimal.NewFromInt(0),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}

		hashedPwd := "$2a$04$Oxfvvna0lq3qToeWXHVj.esXeXleZnTYUBAkhr55JcN73o.J33u7W"

		entityUser := entity.User{
			Fullname: req.Fullname,
			Email:    req.Email,
			Address:  req.Address,
			Password: hashedPwd,
			CityID:   req.CityId,
			Role:     "admin",
		}

		auth.On("HashAndSalt", req.Password).Return(hashedPwd, nil)
		userUsecase.On("IsUserExist", req.Email).Return(&entityUser, false)
		userUsecase.On("CreateUser", entityUser).Return(&entityUser, nil)
		gameUsecase.On("CreateGame", 0).Return(&game, nil)
		walletUsecase.On("CreateWallet", 0).Return(&wallet, nil)
		user, err := uc.Register(req)

		assert.Equal(t, res, user)
		assert.Nil(t, err)
	})

	t.Run("Should return error when failed hash password", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		gameUsecase := new(mocks.GameUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthAdminUsecase(usecase.AuthAdminUsecaseImplementationConfig{
			AuthAdminUsecase: auth,
			UserUsecase:      userUsecase,
			GameUsecase:      gameUsecase,
			WalletUsecase:    walletUsecase,
		})

		req := dto.RegisterRequest{
			Fullname: "Giwang Dwi Kintan",
			Email:    "gidwikintan@gmail.com",
			Password: "password",
			Address:  "Jl. Kebon Jeruk",
			CityId:   1,
		}

		auth.On("HashAndSalt", req.Password).Return("", errors.New("error"))
		user, err := uc.Register(req)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "error")

	})
	// t.Run("Should return success when exist user is successfully registered", func(t *testing.T) {

	// 	userUsecase := new(mocks.UserUsecase)
	// 	walletUsecase := new(mocks.WalletUsecase)
	// 	gameUsecase := new(mocks.GameUsecase)
	// 	auth := new(mocks.AuthUtil)

	// 	uc := usecase.NewAuthAdminUsecase(usecase.AuthAdminUsecaseImplementationConfig{
	// 		AuthAdminUsecase: auth,
	// 		UserUsecase:      userUsecase,
	// 		GameUsecase:      gameUsecase,
	// 		WalletUsecase:    walletUsecase,
	// 	})

	// 	req := dto.RegisterRequest{
	// 		Fullname: "Giwang Dwi Kintan",
	// 		Email:    "gidwikintan@gmail.com",
	// 		Password: "password",
	// 		Address:  "Jl. Kebon Jeruk",
	// 		CityId:   1,
	// 	}
	// 	res := dto.RegisterResponse{
	// 		Fullname: "Giwang Dwi Kintan",
	// 		Email:    "gidwikintan@gmail.com",
	// 		Address:  "Jl. Kebon Jeruk",
	// 		CityId:   1,
	// 		Role:     "admin",
	// 	}

	// 	hashedPwd := "$2a$04$Oxfvvna0lq3qToeWXHVj.esXeXleZnTYUBAkhr55JcN73o.J33u7W"

	// 	entityUser := entity.User{
	// 		Fullname: "Giwang Dwi Kintan",
	// 		Email:    "gidwikintan@gmail.com",
	// 		Password: "password",
	// 		Address:  "Jl. Kebon Jeruk",
	// 		CityID:   1,
	// 		Role:     "admin",
	// 	}

	// 	auth.On("HashAndSalt", req.Password).Return(hashedPwd, nil)
	// 	userUsecase.On("IsUserExist", req.Email).Return(&entityUser, true)
	// 	userUsecase.On("UpdateRole", req.Email, "admin").Return(&entityUser, nil)
	// 	user, err := uc.Register(req)

	// 	assert.Equal(t, res, user)
	// 	assert.Nil(t, err)
	// })

	t.Run("Should return error when failed create user", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		gameUsecase := new(mocks.GameUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthAdminUsecase(usecase.AuthAdminUsecaseImplementationConfig{
			AuthAdminUsecase: auth,
			UserUsecase:      userUsecase,
			GameUsecase:      gameUsecase,
			WalletUsecase:    walletUsecase,
		})

		req := dto.RegisterRequest{
			Fullname: "Giwang Dwi Kintan",
			Email:    "gidwikintan@gmail.com",
			Password: "password",
			Address:  "Jl. Kebon Jeruk",
			CityId:   1,
		}

		hashedPwd := "$2a$04$Oxfvvna0lq3qToeWXHVj.esXeXleZnTYUBAkhr55JcN73o.J33u7W"

		entityUser := entity.User{
			Fullname: req.Fullname,
			Email:    req.Email,
			Address:  req.Address,
			Password: hashedPwd,
			CityID:   req.CityId,
			Role:     "admin",
		}

		auth.On("HashAndSalt", req.Password).Return(hashedPwd, nil)
		userUsecase.On("IsUserExist", req.Email).Return(&entityUser, false)
		userUsecase.On("CreateUser", entityUser).Return(nil, errors.New("error"))
		user, err := uc.Register(req)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "error")
	})
	t.Run("Should return error when failed create game", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		gameUsecase := new(mocks.GameUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthAdminUsecase(usecase.AuthAdminUsecaseImplementationConfig{
			AuthAdminUsecase: auth,
			UserUsecase:      userUsecase,
			GameUsecase:      gameUsecase,
			WalletUsecase:    walletUsecase,
		})

		req := dto.RegisterRequest{
			Fullname: "Giwang Dwi Kintan",
			Email:    "gidwikintan@gmail.com",
			Password: "password",
			Address:  "Jl. Kebon Jeruk",
			CityId:   1,
		}

		hashedPwd := "$2a$04$Oxfvvna0lq3qToeWXHVj.esXeXleZnTYUBAkhr55JcN73o.J33u7W"

		entityUser := entity.User{
			Fullname: req.Fullname,
			Email:    req.Email,
			Address:  req.Address,
			Password: hashedPwd,
			CityID:   req.CityId,
			Role:     "admin",
		}

		auth.On("HashAndSalt", req.Password).Return(hashedPwd, nil)
		userUsecase.On("IsUserExist", req.Email).Return(&entityUser, false)
		userUsecase.On("CreateUser", entityUser).Return(&entityUser, nil)
		gameUsecase.On("CreateGame", 0).Return(nil, errors.New("error"))
		user, err := uc.Register(req)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "error")
	})

	t.Run("Should return error when failed create wallet", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		gameUsecase := new(mocks.GameUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthAdminUsecase(usecase.AuthAdminUsecaseImplementationConfig{
			AuthAdminUsecase: auth,
			UserUsecase:      userUsecase,
			GameUsecase:      gameUsecase,
			WalletUsecase:    walletUsecase,
		})

		req := dto.RegisterRequest{
			Fullname: "Giwang Dwi Kintan",
			Email:    "gidwikintan@gmail.com",
			Password: "password",
			Address:  "Jl. Kebon Jeruk",
			CityId:   1,
		}

		game := entity.Game{
			UserId:           1,
			Chance:           decimal.NewFromInt(0),
			TotalGamesPlayed: decimal.NewFromInt(0),
		}

		hashedPwd := "$2a$04$Oxfvvna0lq3qToeWXHVj.esXeXleZnTYUBAkhr55JcN73o.J33u7W"

		entityUser := entity.User{
			Fullname: req.Fullname,
			Email:    req.Email,
			Address:  req.Address,
			Password: hashedPwd,
			CityID:   req.CityId,
			Role:     "admin",
		}

		auth.On("HashAndSalt", req.Password).Return(hashedPwd, nil)
		userUsecase.On("IsUserExist", req.Email).Return(&entityUser, false)
		userUsecase.On("CreateUser", entityUser).Return(&entityUser, nil)
		gameUsecase.On("CreateGame", 0).Return(&game, nil)
		walletUsecase.On("CreateWallet", 0).Return(nil, errors.New("error"))
		user, err := uc.Register(req)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "error")
	})
}

func TestLoginAdmin(t *testing.T) {
	t.Run("Should return user when succes login", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		gameUsecase := new(mocks.GameUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthAdminUsecase(usecase.AuthAdminUsecaseImplementationConfig{
			AuthAdminUsecase: auth,
			UserUsecase:      userUsecase,
			GameUsecase:      gameUsecase,
			WalletUsecase:    walletUsecase,
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
			Role:     "admin",
		}

		userUsecase.On("GetUserByEmail", req.Email).Return(&userEntity, nil)
		auth.On("ComparePassword", req.Password, userEntity.Password).Return(true)
		auth.On("GenerateAccessToken", &userEntity).Return("token", nil)
		user, err := uc.Login(req)

		assert.Equal(t, res.AccessToken, user.AccessToken)
		assert.Nil(t, err)
	})
	t.Run("Should return error when input invalid email", func(t *testing.T) {
		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		gameUsecase := new(mocks.GameUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthAdminUsecase(usecase.AuthAdminUsecaseImplementationConfig{
			AuthAdminUsecase: auth,
			UserUsecase:      userUsecase,
			GameUsecase:      gameUsecase,
			WalletUsecase:    walletUsecase,
		})
		req := dto.LoginRequest{
			Email:    "gidwikintan@gmail.com",
			Password: "password",
		}
		userUsecase.On("GetUserByEmail", req.Email).Return(nil, errors.New("error"))
		user, err := uc.Login(req)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("Should return error when input invalid password", func(t *testing.T) {

		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		gameUsecase := new(mocks.GameUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthAdminUsecase(usecase.AuthAdminUsecaseImplementationConfig{
			AuthAdminUsecase: auth,
			UserUsecase:      userUsecase,
			GameUsecase:      gameUsecase,
			WalletUsecase:    walletUsecase,
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

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, "Unauthorized", err.Error())
	})
	t.Run("Test case 1: failed generate token", func(t *testing.T) {
		userUsecase := new(mocks.UserUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		gameUsecase := new(mocks.GameUsecase)
		auth := new(mocks.AuthUtil)

		uc := usecase.NewAuthAdminUsecase(usecase.AuthAdminUsecaseImplementationConfig{
			AuthAdminUsecase: auth,
			UserUsecase:      userUsecase,
			GameUsecase:      gameUsecase,
			WalletUsecase:    walletUsecase,
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

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, "Unauthorized", err.Error())
	})

}
