package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Token struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	AccessToken string
	CreatedAt   time.Time
	TokenType   string
}
