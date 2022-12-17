package handler

import "final-project-backend/usecase"

type Handler struct {
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthUsecase
	authAdminUsecase usecase.AuthAdminUsecase
	cityUsecase usecase.CityUsecase
	walletUsecase usecase.WalletUsecase
}

type HandlerConfig struct {
	UserUsecase usecase.UserUsecase
	AuthUsecase usecase.AuthUsecase
	AuthAdminUsecase usecase.AuthAdminUsecase
	CityUsecase usecase.CityUsecase
	WalletUsecase usecase.WalletUsecase
}

func NewHandler(c HandlerConfig) *Handler {
	return &Handler{
		userUsecase: c.UserUsecase,
		authUsecase: c.AuthUsecase,
		authAdminUsecase: c.AuthAdminUsecase,
		cityUsecase: c.CityUsecase,
		walletUsecase: c.WalletUsecase,
	}
}
