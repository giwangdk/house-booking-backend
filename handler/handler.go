package handler

import "final-project-backend/usecase"

type Handler struct {
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthUsecase
}

type HandlerConfig struct {
	UserUsecase usecase.UserUsecase
	AuthUsecase usecase.AuthUsecase
}

func NewHandler(c HandlerConfig) *Handler {
	return &Handler{
		userUsecase: c.UserUsecase,
		authUsecase: c.AuthUsecase,
	}
}


