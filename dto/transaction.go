package dto

import "final-project-backend/entity"

type Transaction struct {
	ID            int         `json:"id"`
	ReservationID int         `json:"reservation_id"`
	UserID        int         `json:"user_id"`
	HouseID       int         `json:"house_id"`
	TransferSlip  string      `json:"transfer_slip"`
	Reservation   Reservation `json:"reservation"`
}

type TransactionList struct {
	Transactions []Transaction `json:"transactions"`
}

func (t *Transaction) BuildResponse(tx entity.Transaction) *Transaction {
	return &Transaction{
		ID:            int(tx.ID),
		ReservationID: tx.ReservationID,
		UserID:        tx.UserID,
		HouseID:       tx.HouseID,
		TransferSlip:  tx.TransferSlip,
		Reservation:   *(&Reservation{}).BuildResponse(tx.Reservation),
	}
}

func (t *TransactionList) BuildResponse(txs []entity.Transaction) *TransactionList {
	var tx []Transaction
	for _, transaction := range txs {
		tx = append(tx, *(&Transaction{}).BuildResponse(transaction))
	}
	return &TransactionList{
		Transactions: tx,
	}
}
