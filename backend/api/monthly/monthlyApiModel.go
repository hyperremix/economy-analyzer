package monthly

import (
	"github.com/hyperremix/economy-analyzer/backend/model"
	"time"
)

type MonthlyApiModel struct {
	ClassifiedTransactions map[model.ClassificationType][]TransactionApiModel
	Month                  time.Time
	LastMonthBalance       float64
}

func ManyNewMonthlyApiModels(monthlies []model.Monthly) (monthlyApiModels []MonthlyApiModel) {
	for _, monthly := range monthlies {
		monthlyApiModels = append(monthlyApiModels, NewMonthlyApiModel(monthly))
	}

	return
}

func NewMonthlyApiModel(monthly model.Monthly) (monthlyApiModel MonthlyApiModel) {
	classifiedTransactions := transformClassifiedTransactions(monthly.ClassifiedTransactions)

	return MonthlyApiModel{
		ClassifiedTransactions: classifiedTransactions,
		Month:            monthly.Month,
		LastMonthBalance: monthly.LastMonthBalance}
}

func transformClassifiedTransactions(classifiedTransactionsMap model.ClassifiedTransactionsMap) map[model.ClassificationType][]TransactionApiModel {
	classifiedTransactions := make(map[model.ClassificationType][]TransactionApiModel)

	for classificationType, transactions := range classifiedTransactionsMap.GetRawMap() {
		classifiedTransactions[classificationType] = ManyNewTransactionApiModels(transactions)
	}

	return classifiedTransactions
}
