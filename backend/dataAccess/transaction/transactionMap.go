package transaction

import (
	"fmt"
	"strconv"
	"time"

	"strings"

	"github.com/hyperremix/economy-analyzer/backend/model"
)

type TransactionMap struct{}

func (transactionMap *TransactionMap) TransformMany(records [][]string) []model.Transaction {

	var transactions = make([]model.Transaction, len(records))

	for i, record := range records {
		transactions[i] = transactionMap.Transform(record)
	}

	return transactions
}

func (transactionMap *TransactionMap) Transform(record []string) model.Transaction {

	parseableFloat := getGolangFloatString(record[7])
	amount, err := strconv.ParseFloat(parseableFloat, 64)

	if err != nil {
		panic(fmt.Sprintf("Unable to parse amount %v for a transaction", parseableFloat))
	}

	t, err := time.Parse("02.01.2006", record[1])

	if err != nil {
		panic(fmt.Sprintf("Unable to transform date of record %v for a transaction", record))
	}

	return model.Transaction{Amount: amount, Client: record[3], Date: t, Purpose: record[4]}
}

func getGolangFloatString(germanFloatString string) string {
	parts := strings.Split(germanFloatString, ",")

	integer := strings.Replace(parts[0], ".", "", -1)

	if len(parts) == 2 {
		decimal := parts[1]
		return integer + "." + decimal
	}

	return integer
}
