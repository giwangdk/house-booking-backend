package server

import (
	"final-project-backend/config"
	"final-project-backend/db"
	"final-project-backend/helper"
	"final-project-backend/repository"
	"final-project-backend/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {

	userRepo := repository.NewPostgresUserRepository(repository.PostgresUserRepositoryConfig{
		DB: db.Get(),
	})

	cityRepo := repository.NewPostgresCityRepository(repository.PostgresCityRepositoryConfig{
		DB: db.Get(),
	})
	walletRepo := repository.NewPostgresWalletRepository(repository.PostgresWalletRepositoryConfig{
		DB: db.Get(),
	})
	gameRepo := repository.NewPostgresGameRepository(repository.PostgresGameRepositoryConfig{
		DB: db.Get(),
	})
	houseRepo := repository.NewPostgresHouseRepository(repository.PostgresHouseRepositoryConfig{
		DB: db.Get(),
	})
	houseDetailRepo := repository.NewPostgresHouseDetailRepository(repository.PostgresHouseDetailRepositoryConfig{
		DB: db.Get(),
	})
	housePhotoRepo := repository.NewPostgresHousePhotoRepository(repository.PostgresHousePhotoRepositoryConfig{
		DB: db.Get(),
	})
	reservationRepo := repository.NewPostgresReservationRepository(repository.PostgresReservationRepositoryConfig{
		DB: db.Get(),
	})
	transactionRepo := repository.NewPostgresTransactionRepository(repository.PostgresTransactionRepositoryConfig{
		DB: db.Get(),
	})
	pickupRepo := repository.NewPostgresPickupRepository(repository.PostgresPickupRepositoryConfig{
		DB: db.Get(),
	})
	pickupStatusRepo := repository.NewPostgresPickupStatusRepository(repository.PostgresPickupStatusRepositoryConfig{
		DB: db.Get(),
	})
	walletTransactionRepo := repository.NewPostgresWalletTransactionRepository(repository.PostgresWalletTransactionRepositoryConfig{
		DB: db.Get(),
	})
	//duration, err := strconv.Atoi(config.Config.AuthConfig.Duration)

	
	auth := helper.NewAuthUtil(helper.AuthUtilImplConfig{
		HmacSampleSecret: config.Config.AuthConfig.HmacSampleSecret,
		//Duration:         *jwt.NewNumericDate(jwt.TimeFunc().Add(time.Duration(1) * time.Minute)),
		//Duration: jwt.NewNumericDate( time.Now().Add(time.Duration(1) * time.Minute)),
	})

	userUsecase := usecase.NewUserUseCase(usecase.UserUsecaseImplementationConfig{
		Repository:  userRepo,
		AuthUsecase: auth,
	})

	walletUsecase := usecase.NewWalletUseCase(usecase.WalletUsecaseImplementationConfig{
		Repository: walletRepo,
	})

	gameUsecase := usecase.NewGameUseCase(usecase.GameUsecaseImplementationConfig{
		Repository: gameRepo,
		WalletUsecase: walletUsecase,
		WalletTxRepo: walletTransactionRepo,
	})

	houseUsecase := usecase.NewHouseUseCase(usecase.HouseUsecaseImplementationConfig{
		Repository: houseRepo,
	})

	houseDetailUsecase := usecase.NewHouseDetailUseCase(usecase.HouseDetailUsecaseImplementationConfig{
		Repository: houseDetailRepo,
	})

	housePhotoUsecase := usecase.NewHousePhotoUseCase(usecase.HousePhotoUsecaseImplementationConfig{
		Repository: housePhotoRepo,
	})

	authUsecase := usecase.NewAuthUsecase(usecase.AuthUsecaseImplementationConfig{
		AuthUsecase:   auth,
		UserUsecase:   userUsecase,
		WalletUsecase: walletUsecase,
		GameUsecase:   gameUsecase,
	})

	authAdminUsecase := usecase.NewAuthAdminUsecase(usecase.AuthAdminUsecaseImplementationConfig{
		AuthAdminUsecase: auth,
		UserUsecase:      userUsecase,
		WalletUsecase:    walletUsecase,
		GameUsecase:      gameUsecase,
	})

	cityUsecase := usecase.NewCityUseCase(usecase.CityUsecaseImplementationConfig{
		Repository: cityRepo,
	})
	pickupUsecase := usecase.NewPickupUseCase(usecase.PickupUsecaseImplementationConfig{
		Repository: pickupRepo,
	})

	pickupStatusUsecase:= usecase.NewPickupStatusUseCase(usecase.PickupStatusUsecaseImplementationConfig{
		Repository: pickupStatusRepo,
	})

	reservationUsecase := usecase.NewReservationUseCase(usecase.ReservationUsecaseImplementationConfig{
		Repository:    reservationRepo,
		UserUsecase:   userUsecase,
		PickupUsecase: pickupUsecase,
	})
	transactionUsecase := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
		Repository:         transactionRepo,
		ReservationUsecase: reservationUsecase,
		HouseUsecase: 	 houseUsecase,
		WalletUsecase: 	 walletUsecase,
		WalletTxRepo: 	 walletTransactionRepo,
	})
	walletTransactionUsecase:= usecase.NewWalletTransactionUseCase(usecase.WalletTransactionUsecaseImplementationConfig{
		Repository: walletTransactionRepo,
		WalletUsecase: walletUsecase,
		GameRepository: gameRepo,
	})

	r := CreateRouter(&RouterConfig{
		AuthUsecase:        authUsecase,
		AuthAdminUsecase:   authAdminUsecase,
		UserUsecase:        userUsecase,
		CityUsecase:        cityUsecase,
		WalletUsecase:      walletUsecase,
		GameUsecase:        gameUsecase,
		HouseUsecase:       houseUsecase,
		HouseDetailUsecase: houseDetailUsecase,
		HousePhotoUsecase:  housePhotoUsecase,
		ReservationUsecase: reservationUsecase,
		TransactionUsecase: transactionUsecase,
		WalletTransactionUsecase: walletTransactionUsecase,
		PickupUsecase : pickupUsecase,
		PickupStatusUsecase: pickupStatusUsecase,
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
