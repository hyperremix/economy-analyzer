package transaction

import (
	"encoding/csv"
	"os"

	"log"

	"github.com/hyperremix/economy-analyzer/model"
)

const repositoryPath string = "C:\\Users\\fredr_000\\goplayground\\src\\github.com\\hyperremix\\economy-analyzer\\transactions.csv"

type TransactionRepository struct{}

func (transactionRepository *TransactionRepository) Find() []model.Transaction {
	fileReader, err := os.Open(repositoryPath)

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

	return new(TransactionMap).TransformMany(records)
}
