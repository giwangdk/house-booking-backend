package server

import (
	"final-project-backend/handler"
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
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	return r
}
