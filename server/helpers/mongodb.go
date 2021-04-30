package helpers

import (
	"log"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func InitDB() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("Local_Database")
}

func Collection() *mgo.Collection {
	return db.C("userProfile")
}
