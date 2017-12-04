package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

/*
MongoDB defines the database to connect to
*/
type MongoDB struct {
	DatabaseURL    string
	DatabaseName   string
	CollectionName string
}

/*
Init connects to DB
*/
func (db *MongoDB) init() {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}
	defer session.Close()
}

/*
Registers a book in the DB
*/
func (db *MongoDB) registerBook(b Book) error {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Inserts book in DB
	err = session.DB(db.DatabaseName).C(db.CollectionName).Insert(b)
	if err != nil {
		fmt.Printf("Error in Insert(): %v", err.Error())
		return err
	}

	return nil
}

/*
Returns the number of books in collection
*/
func (db *MongoDB) CountBooks() int {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	count, err := session.DB(db.DatabaseName).C(db.CollectionName).Count()
	if err != nil {
		fmt.Printf("Error in Count(): %v", err.Error())
		return -1
	}
	// Returns -1 if error, else returns count
	return count
}
