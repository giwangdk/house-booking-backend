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
	GetTransactionsGuest() (*dto.TransactionList, error)
	GetTransactionsUser(userId int) (*dto.TransactionList, error)
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
		return nil, httperror.NotFoundError("Reservation is not found!")
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

	walletRecipient, err := u.walletUsecase.GetWalletByUserID(1)
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
		tx, err := u.repository.GetTransactionByBookingCode(r.BookingCode)
		if err != nil {
			return nil, httperror.BadRequestError("Transaction is not found!", "ERROR_TRANSACTION_NOT_FOUND")
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
		return nil, err
	}

	isValid := u.walletUsecase.IsValidBalance(reservation.TotalPrice, *walletSender)
	if !isValid {
		return nil, httperror.BadRequestError("Insufficient balance!", "ERROR_INSUFFICIENT_BALANCE")
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

	uploadUrl, err := helper.ImageUploadHelper(r.TransferSlip)
	if err != nil {
		return nil, httperror.BadRequestError("Failed to upload image!", "ERROR_UPLOADING_IMAGE")
	}

	transaction, err := u.repository.CreateTransaction(entity.Transaction{
		ReservationID: int(reservation.ID),
		HouseID:       int(house.ID),
		UserID:        reservation.UserID,
		TransferSlip:  uploadUrl,
	})
	if err != nil {
		return nil, httperror.BadRequestError("Failed to create transaction!", "ERROR_CREATING_TRANSACTION")
	}
	_, err = u.reservationUsecase.UpdateStatusReservation(reservation.ID, 4)
	if err != nil {
		return nil, err
	}

	res := (&dto.CreateTransactionResponse{}).BuildResponse(*transaction)

	return res, nil

}

func (u *TransactionUsecaseImplementation) GetTransactionsGuest() (*dto.TransactionList, error) {

	transactions, err := u.repository.GetTransactionsGuest()
	if err != nil {
		return nil, err
	}

	res := (&dto.TransactionList{}).BuildResponse(transactions)

	return res, nil
}

func (u *TransactionUsecaseImplementation) GetTransactionsUser(userId int) (*dto.TransactionList, error) {

	transactions, err := u.repository.GetTransactionsUser(userId)
	if err != nil {
		return nil, err
	}

	res := (&dto.TransactionList{}).BuildResponse(transactions)

	return res, nil
}
