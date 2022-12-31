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
	gameRepository repository.GameRepository
}

type WalletTransactionUsecaseImplementationConfig struct {
	Repository    repository.WalletTransactionRepository
	WalletUsecase WalletUsecase
	GameRepository repository.GameRepository
}

func NewWalletTransactionUseCase(c WalletTransactionUsecaseImplementationConfig) WalletTransactionUsecase {
	return &walletTransactionUsecaseImplementation{
		repository:    c.Repository,
		walletUsecase: c.WalletUsecase,
		gameRepository: c.GameRepository,
	}
}

func (u *walletTransactionUsecaseImplementation) GetWalletTransactionsUser(walletID int) (*[]entity.WalletTransaction, error) {
	Wallettransactions, err := u.repository.GetWalletTransactionsUser(walletID)
	if err != nil {
		return nil, httperror.NotFoundError("Wallet Transaction is not found!")
	}

	return Wallettransactions, nil
}

func (u *walletTransactionUsecaseImplementation) GetWalletTransactions(s string, sortBy string, sort string, limit int, page int) ( *[]entity.WalletTransaction, error) {
	Wallettransactions, err := u.repository.GetWalletTransactions(s, sortBy, sort, limit, page)
	if err != nil {
		return  nil, httperror.NotFoundError("Wallet Transaction is not found!")
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

	game, err:= u.gameRepository.GetGameByUserID(t.Recipient)
	if err != nil {
		return nil, err
	}


	if t.Amount.GreaterThanOrEqual(decimal.NewFromInt(500000)){
		count:= t.Amount.Div(decimal.NewFromInt(500000))
		u.gameRepository.IncreaseChance(int(count.IntPart()), *game)
	} 


	wallet, err := u.walletUsecase.GetWalletByUserID(t.Recipient)
	if err != nil {
		return nil, err
	}


	entity := entity.WalletTransaction{
		Sender:      int64(wallet.ID),
		Recipient:  int64(wallet.ID),
		Amount:      t.Amount,
		Description: "Top Up ",
	}

	
	transaction, err := u.repository.CreateWalletTransaction(entity)
	if err != nil {
		return nil, httperror.BadRequestError("Error creating wallet transaction", "ERROR_CREATING_WALLET_TRANSACTION")
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
