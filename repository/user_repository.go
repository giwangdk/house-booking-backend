package repository

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/httperror"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(user entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUser(userID int) (*entity.User, error)
	EditUser(u entity.User, userId int) (*entity.User, error)
	CreateUserAdmin(u entity.User) (*entity.User, error)
	UpdateRole(u entity.User, role string) (*entity.User, error)
}

type postgresUserRepository struct {
	db *gorm.DB
}

type PostgresUserRepositoryConfig struct {
	DB *gorm.DB
}

func NewPostgresUserRepository(c PostgresUserRepositoryConfig) UserRepository {
	return &postgresUserRepository{
		db: c.DB,
	}
}

func (r *postgresUserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var u entity.User
	err := r.db.Where("email = ?", email).First(&u).Error

	if notFound := errors.Is(err, gorm.ErrRecordNotFound); notFound {
		return nil, httperror.NotFoundError("user not found")
	}
	return &u, nil
}

func (r *postgresUserRepository) CreateUser(u entity.User) (*entity.User, error) {
	res := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoNothing: true,
	}).Create(&u)

	if res.RowsAffected == 0 && res.Error == nil {
		return nil, httperror.BadRequestError("Email already exist", "EMAIL_ALREADY_EXIST")
	}
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_USER")
	}

	return &u, nil
}

func (r *postgresUserRepository) CreateUserAdmin(u entity.User) (*entity.User, error) {
	res := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoNothing: true,
	}).Create(&u)

	if res.RowsAffected == 0 && res.Error == nil {
		err:= r.db.Model(&u).Where("email = ?", u.Email).Update("role", "admin").Error
		if err != nil {
			return nil, httperror.BadRequestError(err.Error(), "ERROR_CREATE_USER_ADMIN")
		}
	}
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_USER")
	}

	return &u, nil
}

func (r *postgresUserRepository) UpdateRole(u entity.User, role string) (*entity.User, error) {
	res := r.db.Model(&u).Where("email = ?", u.Email).Update("role", role)

	if res.RowsAffected == 0 && res.Error == nil {
		return nil, httperror.BadRequestError("Email already exist", "EMAIL_ALREADY_EXIST")
	}
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_USER")
	}

	return &u, nil
}

func (r *postgresUserRepository) GetUser(userID int) (*entity.User, error) {
	var u entity.User
	err := r.db.Where("id = ?", userID).Preload("City").First(&u).Error

	if notFound := errors.Is(err, gorm.ErrRecordNotFound); notFound {
		return nil, httperror.NotFoundError("user not found")
	}
	return &u, nil
}

func (r *postgresUserRepository) EditUser(u entity.User, userId int) (*entity.User, error) {
	res := r.db.Where("id = ?", userId).Updates(&u)

	if res.RowsAffected == 0 && res.Error == nil {
		return nil, httperror.BadRequestError("Email already exist", "EMAIL_ALREADY_EXIST")
	}
	if res.Error != nil {
		return nil, httperror.BadRequestError(res.Error.Error(), "ERROR_CREATE_USER")
	}

	return &u, nil
}
