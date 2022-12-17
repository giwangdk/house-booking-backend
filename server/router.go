package server

import (
	"final-project-backend/handler"
	"final-project-backend/middleware"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	AuthUsecase usecase.AuthUsecase
	UserUsecase usecase.UserUsecase
	CityUsecase usecase.CityUsecase
	WalletUsecase usecase.WalletUsecase
}

func CreateRouter(c *RouterConfig) *gin.Engine {
	h := handler.NewHandler(handler.HandlerConfig{
		UserUsecase: c.UserUsecase,
		AuthUsecase: c.AuthUsecase,
		CityUsecase: c.CityUsecase,
		WalletUsecase: c.WalletUsecase,
	})

	r := gin.Default()

	r.Use(middleware.ApplyCORS())
	r.POST("/login", h.Login)
	r.POST("/register", h.Register)

	r.Use(middleware.Authorize)
	r.GET("/user", h.GetUser)
	r.GET("/cities", h.GetCities)
	r.PUT("/user", h.EditUser)

	r.Use(middleware.IsAdmin)


	return r
}
