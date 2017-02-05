package model

import "time"

type Transaction struct {
	Amount  float64
	Client  string
	Date    time.Time
	Purpose string
}

type ByDate []Transaction

func (a ByDate) Len() int {
	return len(a)
}

func (a ByDate) Less(i, j int) bool {
	return a[i].Date.Before(a[j].Date)
}

func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
