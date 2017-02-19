package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Transaction struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
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
