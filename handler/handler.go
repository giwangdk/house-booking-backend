package handler

import "final-project-backend/usecase"

type Handler struct {
	userUsecase        usecase.UserUsecase
	authUsecase        usecase.AuthUsecase
	authAdminUsecase   usecase.AuthAdminUsecase
	cityUsecase        usecase.CityUsecase
	walletUsecase      usecase.WalletUsecase
	gameUsecase        usecase.GameUsecase
	houseUsecase       usecase.HouseUsecase
	houseDetailUsecase usecase.HouseDetailUsecase
}

type HandlerConfig struct {
	UserUsecase        usecase.UserUsecase
	AuthUsecase        usecase.AuthUsecase
	AuthAdminUsecase   usecase.AuthAdminUsecase
	CityUsecase        usecase.CityUsecase
	WalletUsecase      usecase.WalletUsecase
	GameUsecase        usecase.GameUsecase
	HouseUsecase       usecase.HouseUsecase
	HouseDetailUsecase usecase.HouseDetailUsecase
}

func NewHandler(c HandlerConfig) *Handler {
	return &Handler{
		userUsecase:        c.UserUsecase,
		authUsecase:        c.AuthUsecase,
		authAdminUsecase:   c.AuthAdminUsecase,
		cityUsecase:        c.CityUsecase,
		walletUsecase:      c.WalletUsecase,
		gameUsecase:        c.GameUsecase,
		houseUsecase:       c.HouseUsecase,
		houseDetailUsecase: c.HouseDetailUsecase,
	}
}
