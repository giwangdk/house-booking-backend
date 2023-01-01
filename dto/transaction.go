package dto

import "final-project-backend/entity"

type Transaction struct {
	ID int `json:"id"`
	ReservationID int `json:"reservation_id"`
	UserID int `json:"user_id"`
	HouseID int `json:"house_id"`
	TransferSlip string `json:"transfer_slip"`
	Reservation Reservation `json:"reservation"`
}

type TransactionLits struct{
	Transactions []Transaction `json:"transactions"`


}

func(t *Transaction) BuildResponse(tx entity.Transaction) *Transaction{
	return &Transaction{
		ID: int(tx.ID),
		ReservationID: tx.ReservationID,
		UserID: tx.UserID,
		HouseID: tx.HouseID,
		TransferSlip: tx.TransferSlip,
		Reservation: *(&Reservation{}).BuildResponse(tx.Reservation),
	}
}

func(t *TransactionLits) BuildResponse(txs []entity.Transaction) *TransactionLits{
	var tx []Transaction
	for _, transaction := range txs {
		tx = append(tx, *(&Transaction{}).BuildResponse(transaction))
	}
	return &TransactionLits{
		Transactions: tx,
	}
}