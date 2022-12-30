package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/helper"
	"final-project-backend/httperror"
	"final-project-backend/repository"
)

type TransactionUsecase interface {
	CreateTransaction(r dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error)
	CreateTransactionRequestGuest(r dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error)
}

type TransactionUsecaseImplementation struct {
	repository         repository.TransactionRepository
	reservationUsecase ReservationUsecase
	houseUsecase       HouseUsecase
	walletUsecase      WalletUsecase
	walletTxRepo       repository.WalletTransactionRepository
}

type TransactionUsecaseImplementationConfig struct {
	Repository         repository.TransactionRepository
	ReservationUsecase ReservationUsecase
	HouseUsecase       HouseUsecase
	WalletUsecase      WalletUsecase
	WalletTxRepo       repository.WalletTransactionRepository
}

func NewTransactionUseCase(c TransactionUsecaseImplementationConfig) TransactionUsecase {
	return &TransactionUsecaseImplementation{
		repository:         c.Repository,
		reservationUsecase: c.ReservationUsecase,
		houseUsecase:       c.HouseUsecase,
		walletUsecase:      c.WalletUsecase,
		walletTxRepo:       c.WalletTxRepo,
	}
}

func (u *TransactionUsecaseImplementation) CreateTransaction(r dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error) {

	reservation, err := u.reservationUsecase.GetReservationByBookingCode(r.BookingCode)
	if err != nil {
		return nil, err
	}

	if reservation.StatusID == 3 {
		return nil, httperror.BadRequestError("Reservation has been canceled!", "ERROR_RESERVATION_CANCELED")
	}

	if reservation.StatusID == 2 {
		return nil, httperror.BadRequestError("Reservation has been paid!", "ERROR_RESERVATION_PAID")
	}

	house, err := u.houseUsecase.GetHouseById(reservation.HouseID)
	if err != nil {
		return nil, err
	}

	walletRecipient, err := u.walletUsecase.GetWalletByUserID(16)
	if err != nil {
		return nil, httperror.BadRequestError("Recipient wallet is not found!", "ERROR_GETTING_WALLET")
	}

	if r.IsGuest {
		entity := entity.WalletTransaction{
			Sender:      int64(walletRecipient.ID),
			Recipient:   int64(walletRecipient.ID),
			Amount:      reservation.TotalPrice,
			Description: "Reservation",
		}

		_, err = u.walletTxRepo.CreateWalletTransaction(entity)
		if err != nil {
			return nil, err
		}
		_, err = u.walletUsecase.IncreaseBalance(reservation.TotalPrice, *walletRecipient)
		if err != nil {
			return nil, err
		}
		_, err = u.reservationUsecase.UpdateStatusReservation(reservation.ID, 2)
		if err != nil {
			return nil, err
		}
		tx, err:= u.repository.GetTransactionByBookingCode(r.BookingCode)
		if err != nil {
			return nil, err
		}

		res := (&dto.CreateTransactionResponse{}).BuildResponse(*tx)

		return res, nil
	}

	transaction, err := u.repository.CreateTransaction(entity.Transaction{
		ReservationID: int(reservation.ID),
		HouseID:       int(house.ID),
		UserID:        reservation.UserID,
	})
	if err != nil {
		return nil, err
	}

	walletSender, err := u.walletUsecase.GetWalletByUserID(reservation.UserID)
	if err != nil {
		return nil, httperror.BadRequestError("Recipient wallet is not found!", "ERROR_GETTING_WALLET")
	}

	entity := entity.WalletTransaction{
		Sender:      int64(walletSender.ID),
		Recipient:   int64(walletRecipient.ID),
		Amount:      reservation.TotalPrice,
		Description: "Reservation",
	}

	_, err = u.walletTxRepo.CreateWalletTransaction(entity)
	if err != nil {
		return nil, err
	}

	_, err = u.walletUsecase.IncreaseBalance(reservation.TotalPrice, *walletRecipient)
	if err != nil {
		return nil, err
	}
	_, err = u.walletUsecase.DecreaseBalance(reservation.TotalPrice, *walletSender)
	if err != nil {
		return nil, err
	}

	_, err = u.reservationUsecase.UpdateStatusReservation(reservation.ID, 2)
	if err != nil {
		return nil, err
	}

	res := (&dto.CreateTransactionResponse{}).BuildResponse(*transaction)

	return res, nil

}

func (u *TransactionUsecaseImplementation) CreateTransactionRequestGuest(r dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error) {

	reservation, err := u.reservationUsecase.GetReservationByBookingCode(r.BookingCode)
	if err != nil {
		return nil, err
	}

	if reservation.StatusID == 3 {
		return nil, httperror.BadRequestError("Reservation has been canceled!", "ERROR_RESERVATION_CANCELED")
	}

	if reservation.StatusID == 2 {
		return nil, httperror.BadRequestError("Reservation has been paid!", "ERROR_RESERVATION_PAID")
	}

	if reservation.StatusID == 4 {
		return nil, httperror.BadRequestError("Your payment is waiting for confirmation!", "ERROR_RESERVATION_PAID")
	}

	house, err := u.houseUsecase.GetHouseById(reservation.HouseID)
	if err != nil {
		return nil, err
	}

	uploadUrl, err:= helper.ImageUploadHelper(r.TransferSlip)
	if err != nil {
		return nil, err
	}

	transaction, err := u.repository.CreateTransaction(entity.Transaction{
		ReservationID: int(reservation.ID),
		HouseID:       int(house.ID),
		UserID:        reservation.UserID,
		TransferSlip:  uploadUrl,

	})
	if err != nil {
		return nil, err
	}
	_, err = u.reservationUsecase.UpdateStatusReservation(reservation.ID, 4)
	if err != nil {
		return nil, err
	}

	res := (&dto.CreateTransactionResponse{}).BuildResponse(*transaction)

	return res, nil

}
