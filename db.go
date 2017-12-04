package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
Registers a book in the database
*/
func (db *MongoDB) registerHandler(w http.ResponseWriter, r *http.Request) {
	// Gets JSON data from request body and puts it in object
	decoder := json.NewDecoder(r.Body)

	// Creates new Book to parse JSON into
	var book Book
	book.ID = db.CountBooks() + 1
	err := decoder.Decode(&book)
	if err != nil {
		fmt.Fprintln(w, "Error decoding JSON data: ", err)
	}
	db.registerBook(&book)
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, book)
}

/*
Registers a book in the DB
*/
func (db *MongoDB) registerBook(b *Book) error {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Inserts book in DB
	err = session.DB(db.DatabaseName).C(db.CollectionName).Insert(&b)
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
