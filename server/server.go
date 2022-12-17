package server

import (
	"final-project-backend/config"
	"final-project-backend/db"
	"final-project-backend/helper"
	"final-project-backend/repository"
	"final-project-backend/usecase"
	"fmt"
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
	walletRepo:= repository.NewPostgresWalletRepository(repository.PostgresWalletRepositoryConfig{
		DB:db.Get(),
	})


	//duration, err := strconv.Atoi(config.Config.AuthConfig.Duration)

	auth := helper.NewAuthUtil(helper.AuthUtilImplConfig{
		HmacSampleSecret: config.Config.AuthConfig.HmacSampleSecret,
		Duration:         jwt.NewNumericDate(jwt.TimeFunc().Add(time.Duration(15) * time.Minute)),
	})

	userUsecase := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
		Repository:  userRepo,
		AuthUsecase: auth,
	})

	walletUsecase := usecase.NewWalletUseCase(usecase.WalletUsecaseImplementationConfig{
		Repository: walletRepo,
	})

	authUsecase := usecase.NewAuthUsecase(usecase.AuthUsecaseImplementationConfig{
		AuthUsecase: auth,
		UserUsecase: userUsecase,
		WalletUsecase: walletUsecase,
	})

	authAdminUsecase := usecase.NewAuthAdminUsecase(usecase.AuthAdminUsecaseImplementationConfig{
		AuthAdminUsecase: auth,
		UserUsecase: userUsecase,
		WalletUsecase: walletUsecase,
	})

	cityUsecase := usecase.NewCityUseCase(usecase.CityUsecaseImplementationConfig{
		Repository: cityRepo,
	})

	r := CreateRouter(&RouterConfig{
		AuthUsecase: authUsecase,
		AuthAdminUsecase: authAdminUsecase,
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
