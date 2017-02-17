package dataAccess

import "github.com/hyperremix/economy-analyzer/backend/model"

type TransactionRepository struct{}

func NewTransactionRepository() *TransactionRepository {
	return new(TransactionRepository)
}

func (TransactionRepository *TransactionRepository) FindMany() []model.Transaction {
	var results []model.Transaction

	findMany("transactions", &results)
	return results
}
