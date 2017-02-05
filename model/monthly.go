package model

import "time"

//Monthly stores all the classified transactions for a month
type Monthly struct {
	ClassifiedTransactions map[ClassificationType][]Transaction
	Month                  time.Time
	LastMonthBalance       float64
}
