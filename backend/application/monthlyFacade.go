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

func (mf *MonthlyFacade) Find() []model.Monthly {
	monthlyTransactions := mf.transactionFacade.GetMonthlyTransactions()

	var monthlies []model.Monthly
	var lastMonthly model.Monthly

	for month, transactions := range monthlyTransactions {
		monthly := model.NewMonthly(month, lastMonthly.GetBalance(), mf.GetClassifiedTransactions(transactions))
		monthlies = append(monthlies, monthly)

		lastMonthly = monthly
	}

	return monthlies
}

func (mf *MonthlyFacade) GetClassifiedTransactions(transactions []model.Transaction) model.ClassifiedTransactionsMap {
	classifiedTransactions := model.NewClassifiedTransactionsMap()

	for _, transaction := range transactions {
		classificationType := mf.classificationProvider.Get(transaction)
		classifiedTransactions.Add(classificationType, transaction)
	}

	return classifiedTransactions
}
