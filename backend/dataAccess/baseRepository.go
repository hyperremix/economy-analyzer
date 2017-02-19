package dataAccess

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

const db = "ea"

var session *mgo.Session

func findMany(collectionName string, results interface{}) {
	session := getSession()
	defer session.Close()

	collection := session.DB(db).C(collectionName)

	err := collection.Find(nil).All(results)

	if err != nil {
		panic(err)
	}
}

func getSession() *mgo.Session {
	if session == nil {
		setSession()
	}

	return session.Copy()
}

func setSession() {
	s, err := mgo.Dial("mongodb://localhost")
	s.SetMode(mgo.Monotonic, true)

	mgo.SetLogger(log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile))

	if err != nil {
		panic(err)
	}

	session = s
}
