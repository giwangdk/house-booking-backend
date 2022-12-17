package handler

import "final-project-backend/usecase"

type Handler struct {
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthUsecase
	cityUsecase usecase.CityUsecase
	walletUsecase usecase.WalletUsecase
}

type HandlerConfig struct {
	UserUsecase usecase.UserUsecase
	AuthUsecase usecase.AuthUsecase
	CityUsecase usecase.CityUsecase
	WalletUsecase usecase.WalletUsecase
}

func NewHandler(c HandlerConfig) *Handler {
	return &Handler{
		userUsecase: c.UserUsecase,
		authUsecase: c.AuthUsecase,
		cityUsecase: c.CityUsecase,
		walletUsecase: c.WalletUsecase,
	}
}
