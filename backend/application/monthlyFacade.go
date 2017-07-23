package application

import (
	"github.com/hyperremix/economy-analyzer/backend/model"
)

type MonthlyFacade struct {
	classificationProvider *classificationProvider
	transactionFacade      *transactionFacade
}

func NewMonthlyFacade() *MonthlyFacade {
	return &MonthlyFacade{classificationProvider: NewClassificationProvider(), transactionFacade: NewTransactionFacade()}
}

func (monthlyFacade *MonthlyFacade) Find() []model.Monthly {
	monthlyTransactions := monthlyFacade.transactionFacade.GetMonthlyTransactions()

	var monthlies []model.Monthly
	var lastMonthBalance float64

	for month, transactions := range monthlyTransactions {
		monthly := model.NewMonthly(month, lastMonthBalance)

		for _, transaction := range transactions {
			classificationType := monthlyFacade.classificationProvider.Get(transaction)
			monthly.AddTransaction(classificationType, transaction)
		}

		monthlies = append(monthlies, monthly)
		lastMonthBalance = monthly.GetBalance()
	}

	return monthlies
}
