package model

import (
	"time"
)

//Monthly stores all the classified transactions for a month
type Monthly struct {
	ClassifiedTransactions map[ClassificationType][]Transaction
	Month                  time.Time
	LastMonthBalance       float64
}

func NewMonthly(month time.Time, lastMonthBalance float64) Monthly {
	firstDayOfMonth := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.UTC)
	classifiedTransactions := make(map[ClassificationType][]Transaction)
	return Monthly{ClassifiedTransactions: classifiedTransactions, Month: firstDayOfMonth, LastMonthBalance: lastMonthBalance}
}

func (m Monthly) GetBalance() float64 {
	sum := m.LastMonthBalance

	for _, transactions := range m.ClassifiedTransactions {
		for _, transaction := range transactions {
			sum += transaction.Amount
		}
	}

	return sum
}

func (m Monthly) AddTransaction(classificationType ClassificationType, transaction Transaction) {
	if _, ok := m.ClassifiedTransactions[classificationType]; !ok {
		m.ClassifiedTransactions[classificationType] = []Transaction{transaction}
		return
	}

	m.ClassifiedTransactions[classificationType] = append(m.ClassifiedTransactions[classificationType], transaction)
}
