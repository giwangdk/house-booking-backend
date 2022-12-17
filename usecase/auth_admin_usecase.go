package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/helper"
	"final-project-backend/httperror"
)

type AuthAdminUsecase interface {
	Register(request dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(request dto.LoginRequest) (*dto.LoginResponse, error)
}

type AuthAdminUsecaseImplementation struct {
	authAdminUsecase helper.AuthUtil
	userUsecase UserUsecase
	walletUsecase WalletUsecase
}

type AuthAdminUsecaseImplementationConfig struct {
	AuthAdminUsecase helper.AuthUtil
	UserUsecase UserUsecase
	WalletUsecase WalletUsecase
}

func NewAuthAdminUsecase(a AuthAdminUsecaseImplementationConfig) AuthAdminUsecase {
	return &AuthAdminUsecaseImplementation{
		authAdminUsecase: a.AuthAdminUsecase,
		userUsecase: a.UserUsecase,
		walletUsecase: a.WalletUsecase,
	}
}

func (a *AuthAdminUsecaseImplementation) Register(u dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := a.authAdminUsecase.HashAndSalt(u.Password)
	if err != nil {
		return nil, err
	}

	entityUser := entity.User{
		Fullname: u.Fullname,
		Email:    u.Email,
		Address:  u.Address,
		Password: hashedPassword,
		CityID:   u.CityId,
		Role:     "admin",
	}
	

	userCreated, err := a.userUsecase.CreateUser(entityUser)
	if err != nil {
		return nil, err
	}

	_, err = a.walletUsecase.CreateWallet(int(userCreated.ID))
	if err != nil {
		return nil, err
	}


	user := dto.RegisterResponse{
		Fullname: userCreated.Fullname,
		Email:    userCreated.Email,
		Address:  userCreated.Address,
		City:     userCreated.CityID,
		Role:     userCreated.Role,
	}

	return &user, nil
}

func (a *AuthAdminUsecaseImplementation) Login(u dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := a.userUsecase.GetUserByEmail(u.Email)
	if err != nil {
		return nil, err
	}

	if user.Role != "admin" {
		return nil, httperror.UnauthorizedError()
	}

	isAuth := a.authAdminUsecase.ComparePassword(user.Password, u.Password)
	if !isAuth {
		return nil, httperror.UnauthorizedError()
	}

	token, err := a.authAdminUsecase.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	userResponse := dto.LoginResponse{
		AccessToken: token,
	}

	return &userResponse, nil
}
