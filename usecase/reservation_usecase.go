package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
	"time"
)

type ReservationUsecase interface {
	CreateReservation(r dto.CreateReservationRequest, houseID int) (*dto.CreateReservationResponse, error)
}

type ReservationUsecaseImplementation struct {
	repository repository.ReservationRepository
}

type ReservationUsecaseImplementationConfig struct {
	Repository repository.ReservationRepository
}

func NewReservationUseCase(c ReservationUsecaseImplementationConfig) ReservationUsecase {
	return &ReservationUsecaseImplementation{
		repository: c.Repository,
	}
}


func (u *ReservationUsecaseImplementation) CreateReservation(r dto.CreateReservationRequest, houseID int) (*dto.CreateReservationResponse, error) {

	entityReservation := entity.Reservation{
		CheckIn: r.CheckIn,
		CheckOut: r.CheckOut,
		TotalPrice: r.TotalPrice,
		HouseID: houseID,
		Expired : time.Now().Add(1 * time.Hour),
		StatusID: 1,


	}
	Reservation, err := u.repository.CreateReservation(entityReservation)
	if err != nil {
		return nil, err
	}

	res := (&dto.CreateReservationResponse{}).BuildResponse(*Reservation)

	return res, nil

}
