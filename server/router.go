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
}

func CreateRouter(c *RouterConfig) *gin.Engine {
	h := handler.NewHandler(handler.HandlerConfig{
		UserUsecase: c.UserUsecase,
		AuthUsecase: c.AuthUsecase,
	})

	r := gin.Default()

	r.Use(middleware.ApplyCORS())
	r.POST("/login", h.Login)
	r.POST("/register", h.Register)

	return r
}
