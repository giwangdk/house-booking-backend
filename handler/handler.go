package handler

import "final-project-backend/usecase"

type Handler struct {
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthUsecase
	cityUsecase usecase.CityUsecase
}

type HandlerConfig struct {
	UserUsecase usecase.UserUsecase
	AuthUsecase usecase.AuthUsecase
	CityUsecase usecase.CityUsecase
}

func NewHandler(c HandlerConfig) *Handler {
	return &Handler{
		userUsecase: c.UserUsecase,
		authUsecase: c.AuthUsecase,
		cityUsecase: c.CityUsecase,
	}
}
