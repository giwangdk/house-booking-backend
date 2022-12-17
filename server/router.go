package server

import (
	"final-project-backend/handler"
	"final-project-backend/middleware"
	"final-project-backend/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	AuthUsecase usecase.AuthUsecase
	AuthAdminUsecase usecase.AuthAdminUsecase
	UserUsecase usecase.UserUsecase
	CityUsecase usecase.CityUsecase
	WalletUsecase usecase.WalletUsecase
	GameUsecase usecase.GameUsecase
}

func CreateRouter(c *RouterConfig) *gin.Engine {
	h := handler.NewHandler(handler.HandlerConfig{
		UserUsecase: c.UserUsecase,
		AuthUsecase: c.AuthUsecase,
		AuthAdminUsecase: c.AuthAdminUsecase,
		CityUsecase: c.CityUsecase,
		WalletUsecase: c.WalletUsecase,
		GameUsecase: c.GameUsecase,
	})

	r := gin.Default()

	r.Use(middleware.ApplyCORS())
	r.POST("/login", h.Login)
	r.POST("/admin/login", h.LoginAdmin)
	r.POST("/register", h.Register)
	r.POST("/admin/register", h.RegisterAdmin)

	r.Use(middleware.Authorize)
	r.GET("/user", h.GetUser)
	r.GET("/cities", h.GetCities)
	r.PUT("/user", h.EditUser)
	r.GET("/wallet", h.GetWalletByUserID)

	r.Use(middleware.IsAdmin)


	return r
}
