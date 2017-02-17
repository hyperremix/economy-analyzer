package dataAccess

import (
	"encoding/csv"
	"os"

	"log"

	"github.com/hyperremix/economy-analyzer/backend/model"
)

type TransactionRepository struct {
	transactionMap *transactionMap
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{transactionMap: new(transactionMap)}
}

func (TransactionRepository *TransactionRepository) Find() []model.Transaction {
	fileReader, err := os.Open("C:\\Users\\fredr_000\\goplayground\\src\\github.com\\hyperremix\\economy-analyzer\\transactions.csv")

	if err != nil {
		log.Fatal(err)
		return make([]model.Transaction, 0)
	}

	reader := csv.NewReader(fileReader)
	reader.Comma = ';'

	reader.Read()

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
		return make([]model.Transaction, 0)
	}

	return TransactionRepository.transactionMap.TransformMany(records)
}
