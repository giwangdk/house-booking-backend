package server

import (
	"final-project-backend/handler"
	"final-project-backend/middleware"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	AuthUsecase              usecase.AuthUsecase
	AuthAdminUsecase         usecase.AuthAdminUsecase
	UserUsecase              usecase.UserUsecase
	CityUsecase              usecase.CityUsecase
	WalletUsecase            usecase.WalletUsecase
	GameUsecase              usecase.GameUsecase
	HouseUsecase             usecase.HouseUsecase
	HouseDetailUsecase       usecase.HouseDetailUsecase
	HousePhotoUsecase        usecase.HousePhotoUsecase
	ReservationUsecase       usecase.ReservationUsecase
	TransactionUsecase       usecase.TransactionUsecase
	WalletTransactionUsecase usecase.WalletTransactionUsecase
	PickupUsecase 		  usecase.PickupUsecase
	PickupStatusUsecase 		  usecase.PickupStatusUsecase
}

func CreateRouter(c *RouterConfig) *gin.Engine {
	h := handler.NewHandler(handler.HandlerConfig{
		UserUsecase:        c.UserUsecase,
		AuthUsecase:        c.AuthUsecase,
		AuthAdminUsecase:   c.AuthAdminUsecase,
		CityUsecase:        c.CityUsecase,
		WalletUsecase:      c.WalletUsecase,
		GameUsecase:        c.GameUsecase,
		HouseUsecase:       c.HouseUsecase,
		HouseDetailUsecase: c.HouseDetailUsecase,
		HousePhotoUsecase:  c.HousePhotoUsecase,
		ReservationUsecase: c.ReservationUsecase,
		TransactionUsecase: c.TransactionUsecase,
		WalletTransaction:  c.WalletTransactionUsecase,
		PickupUsecase: c.PickupUsecase,
		PickupStatusUsecase: c.PickupStatusUsecase,
	})

	r := gin.Default()

	r.Use(middleware.ApplyCORS())
	r.POST("/login", h.Login)

	r.POST("/reservation", h.CreateReservation)
	r.GET("/reservation/:id", h.GetReservationByBookingCode)
	r.POST("/transaction", h.CreateTransaction)
	r.POST("/guest/transaction/:booking_code", h.CreateTransactionRequestGuest)
	r.POST("/pickup", h.CreatePickup)
	r.POST("/admin/login", h.LoginAdmin)
	r.POST("/register", h.Register)
	r.POST("/admin/register", h.RegisterAdmin)
	r.GET("/houses", h.GetHouses)
	r.GET("/house/:id", h.GetHouseById)
	r.GET("/cities", h.GetCities)

	r.Use(middleware.Authorize)
	r.GET("/user", h.GetUser)
	r.PUT("/user", h.UpdateUser)
	r.PUT("/change-password", h.ChangePassword)
	r.GET("/wallet", h.GetWalletByUserID)
	r.GET("/game", h.GetGameByUserID)
	r.PUT("/game", h.UpdateGame)
	r.POST("/topup", h.TopUp)
	r.POST("/host", h.CreateHost)
	r.GET("/reservations", h.GetReservationsByUserId)

	r.PUT("/pickup/:id", h.UpdateStatusPickup, middleware.IsAdmin)
	r.GET("/pickups", h.GetPickups, middleware.IsAdmin)
	r.GET("/transactions-guest", h.GetTransactionsGuest, middleware.IsAdmin)
	r.GET("/pickup-status", h.GetPickupStatus, middleware.IsAdmin)
	r.DELETE("/house/:id", h.DeleteHouse, middleware.IsAdminAndHost)
	
	r.Use(middleware.IsHost)
	r.GET("/host/houses", h.GetHousesHost)
	r.POST("/house", h.CreateHouse)
	r.PUT("/house/:id", h.UpdateHouse)
	r.POST("/house-detail/:id", h.CreateHouseDetail)
	r.PUT("/house-detail/:id", h.UpdateHouseDetail)
	r.POST("/house-photo/:id", h.CreateHousePhoto)
	r.DELETE("/house-photo/:id", h.DeleteHousePhoto)

	

	return r
}
