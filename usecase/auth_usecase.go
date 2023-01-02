package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/helper"
	"final-project-backend/httperror"
)

type AuthUsecase interface {
	Register(request dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(request dto.LoginRequest) (*dto.LoginResponse, error)
	
}

type AuthUsecaseImplementation struct {
	authUsecase helper.AuthUtil
	userUsecase UserUsecase
	walletUsecase WalletUsecase
	gameUsecase GameUsecase
}

type AuthUsecaseImplementationConfig struct {
	AuthUsecase helper.AuthUtil
	UserUsecase UserUsecase
	WalletUsecase WalletUsecase
	GameUsecase GameUsecase
}

func NewAuthUsecase(a AuthUsecaseImplementationConfig) AuthUsecase {
	return &AuthUsecaseImplementation{
		authUsecase: a.AuthUsecase,
		userUsecase: a.UserUsecase,
		walletUsecase: a.WalletUsecase,
		gameUsecase: a.GameUsecase,
	}
}



func (a *AuthUsecaseImplementation) Register(u dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := a.authUsecase.HashAndSalt(u.Password)
	if err != nil {
	return nil, err
	}

	entityUser := entity.User{
		Fullname: u.Fullname,
		Email:    u.Email,
		Address:  u.Address,
		Password: hashedPassword,
		CityID:   u.CityId,
		Role:     "user",
	}


	userCreated, err := a.userUsecase.CreateUser(entityUser)
	if err != nil {
		return nil, err
	}

	_, err = a.walletUsecase.CreateWallet(int(userCreated.ID))
	if err != nil {
		return nil, err
	}
	
	_, err = a.gameUsecase.CreateGame(int(userCreated.ID))
	if err != nil {
		return nil, err
	}

	res:= (&dto.RegisterResponse{}).BuildResponse(*userCreated)

	return res, nil
}

func (a *AuthUsecaseImplementation) Login(u dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := a.userUsecase.GetUserByEmail(u.Email)
	if err != nil {
		return nil, err
	}

	isAuth := a.authUsecase.ComparePassword(user.Password, u.Password)
	if !isAuth {
		return nil, httperror.UnauthorizedError()
	}

	token, err := a.authUsecase.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	userResponse := dto.LoginResponse{
		AccessToken: token,
	}

	return &userResponse, nil
}
