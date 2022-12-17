package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/helper"
	"final-project-backend/httperror"
	"final-project-backend/repository"
)

type UserUsecase interface {
	GetUserByEmail(email string) (*entity.User, error)
	IsUserExist(email string) (*entity.User, bool)
	CreateUser(r entity.User) (*entity.User, error)
	GetUser(userID int) (*entity.User, error)
	EditUser(u dto.EditUserRequest, userId int) (*dto.EditUserResponse, error)
	UpdateRole(u entity.User, role string) (*entity.User, error)
}

type userUsecaseImplementation struct {
	repository  repository.UserRepository
	authUsecase helper.AuthUtil
}

type UserUsecaseImplementationConfig struct {
	Repository  repository.UserRepository
	AuthUsecase helper.AuthUtil
}

func NewUserUseCase(c UserUsecaseImplementationConfig) UserUsecase {
	return &userUsecaseImplementation{
		repository:  c.Repository,
		authUsecase: c.AuthUsecase,
	}
}

func (u *userUsecaseImplementation) CreateUser(r entity.User) (*entity.User, error) {
	entityUser := entity.User{
		Fullname: r.Fullname,
		Email:    r.Email,
		Address:  r.Address,
		Password: r.Password,
		CityID:   r.CityID,
		Role:     r.Role,
	}

	if r.Role == "admin" {
		user, err := u.repository.CreateUserAdmin(entityUser)
	if err != nil {
		return nil, err
	}
	return user, nil
	}

	user, err := u.repository.CreateUser(entityUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecaseImplementation) UpdateRole(r entity.User, role string) (*entity.User, error) {
	user, err := u.repository.UpdateRole(r, role)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecaseImplementation)IsUserExist(email string) (*entity.User, bool) {
	user, err := u.repository.GetUserByEmail(email)

	if err != nil {
		return nil, false
	}

	if user.Email == email {
		return user, true
	}

	return nil,false
}

func (u *userUsecaseImplementation) GetUserByEmail(email string) (*entity.User, error) {
	user, err := u.repository.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecaseImplementation) GetUser(userID int) (*entity.User, error) {
	user, err := u.repository.GetUser(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecaseImplementation) EditUser(r dto.EditUserRequest, userId int) (*dto.EditUserResponse, error) {
	user, err := u.GetUser(userId)
	if err != nil {
		return nil, err
	}

	isValid := u.authUsecase.ComparePassword(user.Password, r.OldPassword)
	if !isValid {
		return nil, httperror.BadRequestError("Password is not valid", "BAD_REQUEST")
	}

	hashedPass, err := u.authUsecase.HashAndSalt(r.NewPassword)
	if err != nil {
		return nil, err
	}

	reqUser := entity.User{
		Fullname: r.Fullname,
		Address:  r.Address,
		Password: hashedPass,
	}

	updatedUser, err := u.repository.EditUser(reqUser, userId)
	if err != nil {
		return nil, err
	}

	res := dto.EditUserResponse{
		Fullname: updatedUser.Fullname,
		Address:  updatedUser.Address,
	}

	return &res, nil
}
