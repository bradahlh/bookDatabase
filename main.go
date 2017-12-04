package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Book represents the book type and is sent to DB
type Book struct {
	ID     int
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func main() {
	// Initialize DB (would normally not hardcode user/password)
	bookDB := MongoDB{
		"mongodb://testuser:test123@ds129386.mlab.com:29386/books",
		"books",
		"bookRegistry",
	}

	// Web server
	http.HandleFunc("/register", bookDB.registerHandler)
	http.ListenAndServe(":8080", nil)
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
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, book)
}
