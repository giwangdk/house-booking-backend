package usecase_test

import (
	"errors"
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"final-project-backend/usecase"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	t.Run("Should return success when transaction is created", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     false,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}
		walletSender := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		tx := entity.Transaction{
			ReservationID: int(reservation.ID),
		}
		walletTx := entity.WalletTransaction{
			Sender:      int64(walletSender.ID),
			Amount:      reservation.TotalPrice,
			Recipient:   int64(walletRecipient.ID),
			Description: "Reservation",
		}

		response := dto.CreateTransactionResponse{
			ReservationID: int(reservation.ID),
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("CreateTransaction", tx).Return(&tx, nil)
		walletUsecase.On("GetWalletByUserID", 0).Return(&walletSender, nil)
		walletUsecase.On("IsValidBalance", decimal.NewFromInt(1000000), walletSender).Return(true)
		walletTxRepo.On("CreateWalletTransaction", walletTx).Return(&walletTx, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(1000000), walletRecipient).Return(&walletRecipient, nil)
		walletUsecase.On("DecreaseBalance", decimal.NewFromInt(1000000), walletSender).Return(&walletSender, nil)
		reservationUsecase.On("UpdateStatusReservation", int(reservation.ID), 2).Return(&reservation, nil)
		res, err := uc.CreateTransaction(req)

		assert.Nil(t, err)
		assert.Equal(t, response, *res)
		assert.Equal(t, response.ReservationID, res.ReservationID)
	})
	t.Run("Should return error when reservation by booking code is not found", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     false,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(nil, errors.New("error"))
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Reservation is not found!", err.Error())
	})
	t.Run("Should return error when house is not found!", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     false,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(nil, errors.New("error"))
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("Should return error when wallet is not found", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     false,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(nil, errors.New("error"))
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Recipient wallet is not found!", err.Error())
	})
	t.Run("Should return error when create transaction is fail", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     false,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		tx := entity.Transaction{
			ReservationID: int(reservation.ID),
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("CreateTransaction", tx).Return(nil, errors.New("error"))
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("Should return success when wallet sender is not found", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     false,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		tx := entity.Transaction{
			ReservationID: int(reservation.ID),
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("CreateTransaction", tx).Return(&tx, nil)
		walletUsecase.On("GetWalletByUserID", 0).Return(nil, errors.New("error"))
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("Should return error when balance is Insufficient", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     false,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}
		walletSender := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		tx := entity.Transaction{
			ReservationID: int(reservation.ID),
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("CreateTransaction", tx).Return(&tx, nil)
		walletUsecase.On("GetWalletByUserID", 0).Return(&walletSender, nil)
		walletUsecase.On("IsValidBalance", decimal.NewFromInt(1000000), walletSender).Return(false)
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Insufficient balance!", err.Error())
	})
	t.Run("Should return error when create wallet transaction is fail", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     false,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}
		walletSender := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		tx := entity.Transaction{
			ReservationID: int(reservation.ID),
		}
		walletTx := entity.WalletTransaction{
			Sender:      int64(walletSender.ID),
			Amount:      reservation.TotalPrice,
			Recipient:   int64(walletRecipient.ID),
			Description: "Reservation",
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("CreateTransaction", tx).Return(&tx, nil)
		walletUsecase.On("GetWalletByUserID", 0).Return(&walletSender, nil)
		walletUsecase.On("IsValidBalance", decimal.NewFromInt(1000000), walletSender).Return(true)
		walletTxRepo.On("CreateWalletTransaction", walletTx).Return(nil, errors.New("error"))
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("Should return error when fail increase balance", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     false,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}
		walletSender := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		tx := entity.Transaction{
			ReservationID: int(reservation.ID),
		}
		walletTx := entity.WalletTransaction{
			Sender:      int64(walletSender.ID),
			Amount:      reservation.TotalPrice,
			Recipient:   int64(walletRecipient.ID),
			Description: "Reservation",
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("CreateTransaction", tx).Return(&tx, nil)
		walletUsecase.On("GetWalletByUserID", 0).Return(&walletSender, nil)
		walletUsecase.On("IsValidBalance", decimal.NewFromInt(1000000), walletSender).Return(true)
		walletTxRepo.On("CreateWalletTransaction", walletTx).Return(&walletTx, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(1000000), walletRecipient).Return(nil, errors.New("error"))
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("Should return error when fail decrease balance ", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     false,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}
		walletSender := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		tx := entity.Transaction{
			ReservationID: int(reservation.ID),
		}
		walletTx := entity.WalletTransaction{
			Sender:      int64(walletSender.ID),
			Amount:      reservation.TotalPrice,
			Recipient:   int64(walletRecipient.ID),
			Description: "Reservation",
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("CreateTransaction", tx).Return(&tx, nil)
		walletUsecase.On("GetWalletByUserID", 0).Return(&walletSender, nil)
		walletUsecase.On("IsValidBalance", decimal.NewFromInt(1000000), walletSender).Return(true)
		walletTxRepo.On("CreateWalletTransaction", walletTx).Return(&walletTx, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(1000000), walletRecipient).Return(&walletRecipient, nil)
		walletUsecase.On("DecreaseBalance", decimal.NewFromInt(1000000), walletSender).Return(nil, errors.New("error"))

		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("Should return error when  fail update status reservation", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     false,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}
		walletSender := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		tx := entity.Transaction{
			ReservationID: int(reservation.ID),
		}
		walletTx := entity.WalletTransaction{
			Sender:      int64(walletSender.ID),
			Amount:      reservation.TotalPrice,
			Recipient:   int64(walletRecipient.ID),
			Description: "Reservation",
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("CreateTransaction", tx).Return(&tx, nil)
		walletUsecase.On("GetWalletByUserID", 0).Return(&walletSender, nil)
		walletUsecase.On("IsValidBalance", decimal.NewFromInt(1000000), walletSender).Return(true)
		walletTxRepo.On("CreateWalletTransaction", walletTx).Return(&walletTx, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(1000000), walletRecipient).Return(&walletRecipient, nil)
		walletUsecase.On("DecreaseBalance", decimal.NewFromInt(1000000), walletSender).Return(&walletSender, nil)
		reservationUsecase.On("UpdateStatusReservation", int(reservation.ID), 2).Return(nil, errors.New("error"))
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
}

func TestCreateTransactionGuest(t *testing.T) {
	t.Run("Should return success when transaction is created", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     true,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		tx := entity.Transaction{
			ReservationID: int(reservation.ID),
		}
		walletTx := entity.WalletTransaction{
			Sender:      int64(walletRecipient.ID),
			Amount:      reservation.TotalPrice,
			Recipient:   int64(walletRecipient.ID),
			Description: "Reservation",
		}

		response := dto.CreateTransactionResponse{
			ReservationID: int(reservation.ID),
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("GetTransactionByBookingCode", req.BookingCode).Return(&tx, nil)
		walletTxRepo.On("CreateWalletTransaction", walletTx).Return(&walletTx, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(1000000), walletRecipient).Return(&walletRecipient, nil)
		reservationUsecase.On("UpdateStatusReservation", int(reservation.ID), 2).Return(&reservation, nil)
		res, err := uc.CreateTransaction(req)

		assert.Nil(t, err)
		assert.Equal(t, response, *res)
		assert.Equal(t, response.ReservationID, res.ReservationID)
	})

	t.Run("Should return error when transaction by booking code  is not found", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     true,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("GetTransactionByBookingCode", req.BookingCode).Return(nil, errors.New("error"))
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Transaction is not found!", err.Error())
	})
	t.Run("Should return error when failed create wallet transaction", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     true,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		tx := entity.Transaction{
			ReservationID: int(reservation.ID),
		}
		walletTx := entity.WalletTransaction{
			Sender:      int64(walletRecipient.ID),
			Amount:      reservation.TotalPrice,
			Recipient:   int64(walletRecipient.ID),
			Description: "Reservation",
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("GetTransactionByBookingCode", req.BookingCode).Return(&tx, nil)
		walletTxRepo.On("CreateWalletTransaction", walletTx).Return(nil, errors.New("error"))
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})

	t.Run("Should return error when fail increase balance", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     true,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		tx := entity.Transaction{
			ReservationID: int(reservation.ID),
		}
		walletTx := entity.WalletTransaction{
			Sender:      int64(walletRecipient.ID),
			Amount:      reservation.TotalPrice,
			Recipient:   int64(walletRecipient.ID),
			Description: "Reservation",
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("GetTransactionByBookingCode", req.BookingCode).Return(&tx, nil)
		walletTxRepo.On("CreateWalletTransaction", walletTx).Return(&walletTx, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(1000000), walletRecipient).Return(nil, errors.New("error"))
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})
	t.Run("Should return success when fail update reservation status", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		req := dto.CreateTransactionRequest{
			BookingCode: "123",
			IsGuest:     true,
		}

		reservation := entity.Reservation{
			BookingCode: "123",
			TotalPrice:  decimal.NewFromInt(1000000),
		}

		houseDetail := dto.House{
			Name: "House 1",
		}

		walletRecipient := entity.Wallet{
			Balance: decimal.NewFromInt(1000000),
		}

		tx := entity.Transaction{
			ReservationID: int(reservation.ID),
		}
		walletTx := entity.WalletTransaction{
			Sender:      int64(walletRecipient.ID),
			Amount:      reservation.TotalPrice,
			Recipient:   int64(walletRecipient.ID),
			Description: "Reservation",
		}

		reservationDetail := dto.ReservationDetail{
			BookingCode: reservation.BookingCode,
			TotalPrice:  reservation.TotalPrice,
		}

		reservationUsecase.On("GetReservationByBookingCode", req.BookingCode).Return(&reservationDetail, nil)
		houseUsecase.On("GetHouseById", reservation.HouseID).Return(&houseDetail, nil)
		walletUsecase.On("GetWalletByUserID", 1).Return(&walletRecipient, nil)
		repo.On("GetTransactionByBookingCode", req.BookingCode).Return(&tx, nil)
		walletTxRepo.On("CreateWalletTransaction", walletTx).Return(&walletTx, nil)
		walletUsecase.On("IncreaseBalance", decimal.NewFromInt(1000000), walletRecipient).Return(&walletRecipient, nil)
		reservationUsecase.On("UpdateStatusReservation", int(reservation.ID), 2).Return(nil, errors.New("error"))
		res, err := uc.CreateTransaction(req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())
	})

}

func TestGetTransactionGuest(t *testing.T) {
	t.Run("Should return success when get transaction guest", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		var transactions []entity.Transaction

		var response dto.TransactionList

		repo.On("GetTransactionsGuest").Return(transactions, nil)

		res, err := uc.GetTransactionsGuest()

		assert.Nil(t, err)
		assert.Equal(t, &response, res)

	})
	t.Run("Should return error when fail get transaction guest", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		repo.On("GetTransactionsGuest").Return(nil, errors.New("error"))

		res, err := uc.GetTransactionsGuest()

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())

	})
}

func TestGetTransactionUser(t *testing.T) {
	t.Run("Should return success when get transaction guest", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		var transactions []entity.Transaction

		var response dto.TransactionList

		repo.On("GetTransactionsUser", 1).Return(transactions, nil)

		res, err := uc.GetTransactionsUser(1)

		assert.Nil(t, err)
		assert.Equal(t, &response, res)

	})
	t.Run("Should return error when fail get transaction guest", func(t *testing.T) {
		repo := new(mocks.TransactionRepository)
		reservationUsecase := new(mocks.ReservationUsecase)
		houseUsecase := new(mocks.HouseUsecase)
		walletUsecase := new(mocks.WalletUsecase)
		walletTxRepo := new(mocks.WalletTransactionRepository)

		uc := usecase.NewTransactionUseCase(usecase.TransactionUsecaseImplementationConfig{
			Repository:         repo,
			ReservationUsecase: reservationUsecase,
			HouseUsecase:       houseUsecase,
			WalletUsecase:      walletUsecase,
			WalletTxRepo:       walletTxRepo,
		})

		repo.On("GetTransactionsUser", 1).Return(nil, errors.New("error"))

		res, err := uc.GetTransactionsUser(1)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error", err.Error())

	})
}
