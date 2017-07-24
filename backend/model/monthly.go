package model

import (
	"time"
)

//Monthly stores all the classified transactions for a month
type Monthly struct {
	ClassifiedTransactions ClassifiedTransactionsMap
	Month                  time.Time
	LastMonthBalance       float64
}

func NewMonthly(month time.Time, lastMonthBalance float64, classifiedTransactions ClassifiedTransactionsMap) Monthly {
	firstDayOfMonth := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.UTC)
	return Monthly{ClassifiedTransactions: classifiedTransactions, Month: firstDayOfMonth, LastMonthBalance: lastMonthBalance}
}

func (m Monthly) GetBalance() float64 {
	return m.LastMonthBalance + m.ClassifiedTransactions.Sum()
}
