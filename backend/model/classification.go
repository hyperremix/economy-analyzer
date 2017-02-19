package model

import "gopkg.in/mgo.v2/bson"

//Classification determines the classification of a transaction
type Classification struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Client  string
	Purpose string
	Type    ClassificationType
}
