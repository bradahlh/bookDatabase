package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

/*
DB defines the database to connect to
*/
type DB struct {
	DatabaseURL    string
	DatabaseName   string
	CollectionName string
}

/*
Init connects to DB
*/
func (db *DB) Init() {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}
	defer session.Close()
}
