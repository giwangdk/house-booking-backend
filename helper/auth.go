package helper

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthUtil interface {
	GenerateAccessToken(user *entity.User) (string, error)
	ComparePassword(hashedPwd string, inputPwd string) bool
	HashAndSalt(pwd string) (string, error)
}

type authUtilImpl struct {
	hmacSampleSecret string
	duration         jwt.NumericDate
}

type AuthUtilImplConfig struct {
	HmacSampleSecret string
	Duration         jwt.NumericDate
}

func NewAuthUtil(a AuthUtilImplConfig) AuthUtil {
	return &authUtilImpl{
		hmacSampleSecret: a.HmacSampleSecret,
		duration:         *jwt.NewNumericDate(jwt.TimeFunc().Add(time.Duration(1) * time.Minute)),
	}
}

type customClaims struct {
	jwt.RegisteredClaims
	User *dto.UserJWT `json:"user"`
}

func (a *authUtilImpl) HashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (a *authUtilImpl) GenerateAccessToken(user *entity.User) (string, error) {
	claims := &customClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(jwt.TimeFunc().Add(time.Duration(15) * time.Minute)),
			Issuer:    "assignment-golang-backend",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: &dto.UserJWT{
			ID:   int(user.ID),
			Role: user.Role,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.hmacSampleSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *authUtilImpl) ComparePassword(hashedPwd string, inputPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
	return err == nil
}
