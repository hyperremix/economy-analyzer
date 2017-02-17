package application

import (
	"sort"
	"time"

	"github.com/hyperremix/economy-analyzer/backend/dataAccess"
	"github.com/hyperremix/economy-analyzer/backend/model"
)

type MonthlyFacade struct {
	classificationProvider   *classificationProvider
	classificationRepository *dataAccess.ClassificationRepository
	transactionRepository    *dataAccess.TransactionRepository
}

func NewMonthlyFacade() *MonthlyFacade {
	return &MonthlyFacade{classificationProvider: new(classificationProvider), classificationRepository: dataAccess.NewClassificationRepository(), transactionRepository: dataAccess.NewTransactionRepository()}
}

func (monthlyFacade *MonthlyFacade) Find() []model.Monthly {
	classifications := monthlyFacade.classificationRepository.FindMany()
	transactions := monthlyFacade.transactionRepository.FindMany()
	sort.Sort(model.ByDate(transactions))

	currentDate := transactions[0].Date
	var monthlies []model.Monthly
	var classifiedTransactions = make(map[model.ClassificationType][]model.Transaction)
	var lastMonthBalance float64

	for _, transaction := range transactions {
		transactionDate := transaction.Date

		if transactionDate.Month() != currentDate.Month() {
			date := time.Date(currentDate.Year(), currentDate.Month(), 1, 0, 0, 0, 0, time.UTC)
			monthlies = append(monthlies, model.Monthly{Month: date, ClassifiedTransactions: classifiedTransactions, LastMonthBalance: lastMonthBalance})

			lastMonthBalance = getLastMonthBalance(classifiedTransactions, lastMonthBalance)
			classifiedTransactions = make(map[model.ClassificationType][]model.Transaction)
		}

		currentDate = transactionDate
		transactionType := monthlyFacade.classificationProvider.Get(transaction, classifications)

		if _, ok := classifiedTransactions[transactionType]; !ok {
			classifiedTransactions[transactionType] = []model.Transaction{transaction}
			continue
		}

		classifiedTransactions[transactionType] = append(classifiedTransactions[transactionType], transaction)
	}

	return monthlies
}

func getLastMonthBalance(classifiedTransactions map[model.ClassificationType][]model.Transaction, lastMonthBalance float64) float64 {
	sum := lastMonthBalance

	for _, transactions := range classifiedTransactions {
		for _, transaction := range transactions {
			sum += transaction.Amount
		}
	}

	return sum
}
