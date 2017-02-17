package dataAccess

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const db = "ea"

var session *mgo.Session

func findMany(collectionName string, results interface{}) {
	session := getSession()
	defer session.Close()

	collection := session.DB(db).C(collectionName)

	err := collection.Find(bson.M{}).All(results)

	if err != nil {
		panic(err)
	}
}

func getSession() *mgo.Session {
	if session != nil {
		setSession()
	}

	return session.Copy()
}

func setSession() {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	session = s
}
