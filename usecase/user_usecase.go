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
	GetUser(userID int) (*dto.UserDetail, error)
	UpdateUser(u dto.UpdateUserRequest, userId int) (*dto.UpdateUserResponse, error)
	UpdateRole(email string, role string) (*entity.User, error)
	IsUserExistByEmail(email string) (*entity.User, bool)
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


	if r.Role == "admin" {
		user, err := u.repository.CreateUserAdmin(r)
	if err != nil {
		return nil, err
	}

	return user, nil
	}

	user, err := u.repository.CreateUser(r)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecaseImplementation) UpdateRole(email string, role string) (*entity.User, error) {
	user, err := u.repository.UpdateRole(email, role)
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

func (u *userUsecaseImplementation) GetUser(userID int) (*dto.UserDetail, error) {
	user, err := u.repository.GetUser(userID)

	if err != nil {
		return nil, err
	}

	res:= (&dto.UserDetail{}).BuildResponse(*user)

	return res, nil
}

func (u *userUsecaseImplementation) UpdateUser(r dto.UpdateUserRequest, userId int) (*dto.UpdateUserResponse, error) {
	user, err := u.repository.GetUser(userId)
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

	updatedUser, err := u.repository.UpdateUser(reqUser, userId)
	if err != nil {
		return nil, err
	}

	res := (&dto.UpdateUserResponse{}).BuildResponse(*updatedUser)

	return res, nil
}


func (u *userUsecaseImplementation) IsUserExistByEmail(email string) (*entity.User, bool) {
	user, err := u.repository.GetUserByEmail(email)

	if err != nil {
		return nil, false
	}


	return user,true
}