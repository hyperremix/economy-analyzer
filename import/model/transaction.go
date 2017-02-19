package model

import "time"

type Transaction struct {
	Amount  float64
	Client  string
	Date    time.Time
	Purpose string
}
