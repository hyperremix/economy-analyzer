package application

import (
	"github.com/hyperremix/economy-analyzer/backend/dataAccess"
	"github.com/hyperremix/economy-analyzer/backend/model"
	"time"
)

type transactionFacade struct {
	transactionRepository *dataAccess.TransactionRepository
}

func NewTransactionFacade() *transactionFacade {
	return &transactionFacade{transactionRepository: dataAccess.NewTransactionRepository()}
}

func (transactionFacade *transactionFacade) GetMonthlyTransactions() map[time.Time][]model.Transaction {
	transactions := transactionFacade.transactionRepository.FindMany()

	monthlyTransactions := make(map[time.Time][]model.Transaction)

	for _, transaction := range transactions {
		AddTransaction(monthlyTransactions, transaction)
	}

	return monthlyTransactions
}

func AddTransaction(monthlyTransactions map[time.Time][]model.Transaction, transaction model.Transaction) {
	transactionMonth := transaction.GetTransactionMonth()
	if _, ok := monthlyTransactions[transactionMonth]; !ok {
		monthlyTransactions[transactionMonth] = []model.Transaction{transaction}
		return
	}

	monthlyTransactions[transactionMonth] = append(monthlyTransactions[transactionMonth], transaction)
}
