package server

import (
	"final-project-backend/config"
	"final-project-backend/db"
	"final-project-backend/helper"
	"final-project-backend/repository"
	"final-project-backend/usecase"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func initRouter() *gin.Engine {

	userRepo := repository.NewPostgresUserRepository(repository.PostgresUserRepositoryConfig{
		DB: db.Get(),
	})

	cityRepo := repository.NewPostgresCityRepository(repository.PostgresCityRepositoryConfig{
		DB: db.Get(),
	})

	duration, err := strconv.Atoi(config.Config.AuthConfig.Duration)
	if err != nil {
		fmt.Println("error while parsing duration", err)
		return nil
	}

	auth := helper.NewAuthUtil(helper.AuthUtilImplConfig{
		HmacSampleSecret: config.Config.AuthConfig.HmacSampleSecret,
		Duration:         jwt.NewNumericDate(jwt.TimeFunc().Add(time.Duration(duration) * time.Minute)),
	})

	userUsecase := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
		Repository:  userRepo,
		AuthUsecase: auth,
	})

	authUsecase := usecase.NewAuthUsecase(usecase.AuthUsecaseImplementationConfig{
		AuthUsecase: auth,
		UserUsecase: userUsecase,
	})

	cityUsecase := usecase.NewCityUseCase(usecase.CityUsecaseImplementationConfig{
		Repository: cityRepo,
	})

	r := CreateRouter(&RouterConfig{
		AuthUsecase: authUsecase,
		UserUsecase: userUsecase,
		CityUsecase: cityUsecase,
	})
	return r
}

func Init() {
	r := initRouter()
	err := r.Run()
	if err != nil {
		fmt.Println("error while running server", err)
		return
	}

}
