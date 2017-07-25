package monthly

import (
	"github.com/hyperremix/economy-analyzer/backend/model"
	"time"
)

type TransactionApiModel struct {
	ID      string 	
	Amount  float64
	Client  string
	Date    time.Time
	Purpose string
}

func ManyNewTransactionApiModels(transactions []model.Transaction) (transactionApiModels []TransactionApiModel) {
	for _, transaction := range transactions {
		transactionApiModels = append(transactionApiModels, NewTransactionApiModel(transaction))
	}

	return
}

func NewTransactionApiModel(transaction model.Transaction) (transactionApiModel TransactionApiModel) {
	return TransactionApiModel{
		ID:      transaction.ID.Hex(),
		Amount:  transaction.Amount,
		Client:  transaction.Client,
		Date:    transaction.Date,
		Purpose: transaction.Purpose}
}
