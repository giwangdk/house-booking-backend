package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/httperror"
	"final-project-backend/repository"

	"github.com/shopspring/decimal"
)

type WalletTransactionUsecase interface {
	GetWalletTransactionsUser(walletID int) (*[]entity.WalletTransaction, error)
	TopUp(t dto.TopUpRequest) (*dto.TopUpResponse, error)
	GetWalletTransactions(s string, sortBy string, sort string, limit int, page int) ( *[]entity.WalletTransaction, error)
}

type walletTransactionUsecaseImplementation struct {
	repository    repository.WalletTransactionRepository
	walletUsecase WalletUsecase
}

type WalletTransactionUsecaseImplementationConfig struct {
	Repository    repository.WalletTransactionRepository
	WalletUsecase WalletUsecase
}

func NewWalletTransactionUseCase(c WalletTransactionUsecaseImplementationConfig) WalletTransactionUsecase {
	return &walletTransactionUsecaseImplementation{
		repository:    c.Repository,
		walletUsecase: c.WalletUsecase,
	}
}

func (u *walletTransactionUsecaseImplementation) GetWalletTransactionsUser(walletID int) (*[]entity.WalletTransaction, error) {
	Wallettransactions, err := u.repository.GetWalletTransactionsUser(walletID)

	if err != nil {
		return nil, err
	}

	return Wallettransactions, nil
}

func (u *walletTransactionUsecaseImplementation) GetWalletTransactions(s string, sortBy string, sort string, limit int, page int) ( *[]entity.WalletTransaction, error) {
	Wallettransactions, err := u.repository.GetWalletTransactions(s, sortBy, sort, limit, page)
	if err != nil {
		return  nil, err
	}


	return  Wallettransactions, nil
}

func (u *walletTransactionUsecaseImplementation) isValidAmountTopUp(amount decimal.Decimal) bool {
	return amount.GreaterThanOrEqual(decimal.NewFromInt(50000)) && amount.LessThanOrEqual(decimal.NewFromInt(100000000))
}

func (u *walletTransactionUsecaseImplementation) TopUp(t dto.TopUpRequest) (*dto.TopUpResponse, error) {
	if !u.isValidAmountTopUp(t.Amount) {
		return nil, httperror.BadRequestError("Invalid amount minimum Rp.50.000 max Rp. 10 million", "INVALID_AMOUNT")
	}


	wallet, err := u.walletUsecase.GetWalletByUserID(t.Recipient)
	if err != nil {
		return nil, httperror.BadRequestError("Recipient wallet is not found!", "ERROR_GETTING_WALLET")
	}


	entity := entity.WalletTransaction{
		Sender:      int64(wallet.ID),
		Recipient:  int64(wallet.ID),
		Amount:      t.Amount,
		Description: "Top Up ",
	}

	
	transaction, err := u.repository.CreateWalletTransaction(entity)
	if err != nil {
		return nil, err
	}

	_, err = u.walletUsecase.IncreaseBalance(t.Amount, *wallet)
	if err != nil {
		return nil, err
	}

	res := dto.TopUpResponse{
		Sender:      int(transaction.Sender),
		Amount:      transaction.Amount,
		Description: transaction.Description,
		Recipient:   int(transaction.Recipient),
	}

	return &res, nil

}
