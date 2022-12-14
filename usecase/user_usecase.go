package usecase

import (
	"final-project-backend/entity"
	"final-project-backend/repository"
)

type UserUsecase interface {
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(r entity.User) (*entity.User, error)
}

type userUsecaseImplementation struct {
	repository repository.UserRepository
}

type UserUsecaseImplementationConfig struct {
	Repository repository.UserRepository
}

func NewUserUseCase(c UserUsecaseImplementationConfig) UserUsecase {
	return &userUsecaseImplementation{
		repository: c.Repository,
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
	user, err := u.repository.CreateUser(entityUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecaseImplementation) GetUserByEmail(email string) (*entity.User, error) {
	user, err := u.repository.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
