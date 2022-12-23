package server

import (
	"final-project-backend/handler"
	"final-project-backend/middleware"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	AuthUsecase        usecase.AuthUsecase
	AuthAdminUsecase   usecase.AuthAdminUsecase
	UserUsecase        usecase.UserUsecase
	CityUsecase        usecase.CityUsecase
	WalletUsecase      usecase.WalletUsecase
	GameUsecase        usecase.GameUsecase
	HouseUsecase       usecase.HouseUsecase
	HouseDetailUsecase usecase.HouseDetailUsecase
	HousePhotoUsecase  usecase.HousePhotoUsecase
	ReservationUsecase usecase.ReservationUsecase
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
	})

	r := gin.Default()

	r.Use(middleware.ApplyCORS())
	r.POST("/login", h.Login)

	r.POST("/reservation", h.CreateReservation)
	r.POST("/admin/login", h.LoginAdmin)
	r.POST("/register", h.Register)
	r.POST("/admin/register", h.RegisterAdmin)
	r.GET("/houses", h.GetHouses)
	r.GET("/house/:id", h.GetHouseById)
	r.GET("/cities", h.GetCities)

	r.Use(middleware.Authorize)
	r.GET("/user", h.GetUser)
	r.PUT("/user", h.UpdateUser)
	r.GET("/wallet", h.GetWalletByUserID)
	r.GET("/game", h.GetGameByUserID)

	//	r.Use(middleware.IsHost)
	r.POST("/house", h.CreateHouse)
	r.PUT("/house/:id", h.UpdateHouse)
	r.POST("/house-detail/:id", h.CreateHouseDetail)
	r.PUT("/house-detail/:id", h.UpdateHouseDetail)
	r.POST("/house-photo/:id", h.CreateHousePhoto)
	r.DELETE("/house-photo/:id", h.DeleteHousePhoto)

	r.Use(middleware.IsAdmin)

	return r
}
