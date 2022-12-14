package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/helper"
)

type AuthUsecase interface {
	Register(request dto.RegisterRequest) (*dto.RegisterResponse, error)
}

type AuthUsecaseImplementation struct {
	authUsecase helper.AuthUtil
	userUsecase UserUsecase
}

type AuthUsecaseImplementationConfig struct {
	AuthUsecase helper.AuthUtil
	UserUsecase UserUsecase
}

func NewAuthUsecase(a AuthUsecaseImplementationConfig) AuthUsecase {
	return &AuthUsecaseImplementation{
		authUsecase: a.AuthUsecase,
		userUsecase: a.UserUsecase,
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

	user := dto.RegisterResponse{
		Fullname: userCreated.Fullname,
		Email:    userCreated.Email,
		Address:  userCreated.Address,
		City:     userCreated.CityID,
		Role:     userCreated.Role,
	}

	return &user, nil

}
