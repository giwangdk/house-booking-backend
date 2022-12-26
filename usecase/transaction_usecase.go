package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
)

type TransactionUsecase interface {
	CreateTransaction(r dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error)
}

type TransactionUsecaseImplementation struct {
	repository repository.TransactionRepository
	reservationUsecase ReservationUsecase
}

type TransactionUsecaseImplementationConfig struct {
	Repository repository.TransactionRepository
	ReservationUsecase ReservationUsecase
}

func NewTransactionUseCase(c TransactionUsecaseImplementationConfig) TransactionUsecase {
	return &TransactionUsecaseImplementation{
		repository: c.Repository,
		reservationUsecase: c.ReservationUsecase,
	}
}


func (u *TransactionUsecaseImplementation) CreateTransaction (r dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error) {

	reservation, err := u.reservationUsecase.GetReservationByBookingCode(r.BookingCode)
	if err != nil {
		return nil, err
	}

	transaction, err := u.repository.CreateTransaction(entity.Transaction{
		ReservationID: int(reservation.ID),
		HouseID: reservation.HouseID,
		UserID: reservation.UserID,
	})
	if err != nil {
		return nil, err
	}

	_, err = u.reservationUsecase.UpdateStatusReservation(r.ReservationID, 2)
	if err != nil {
		return nil, err
	}

	res := (&dto.CreateTransactionResponse{}).BuildResponse(*transaction)

	return res, nil
	
}