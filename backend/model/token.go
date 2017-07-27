package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Token struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	UserID      bson.ObjectId
	AccessToken string
	CreatedAt   time.Time
	ExpiresIn   float64
}
