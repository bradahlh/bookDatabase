package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// Book represents the book type and is sent to DB
type Book struct {
	ID     bson.ObjectId `bson:"_id, omitempty"`
	Title  string        `json:"title"`
	Author string        `json:"author"`
	Year   int           `json:"year"`
}

func main() {
	// Web server
	http.HandleFunc("/register", registerHandler)
	http.ListenAndServe(":8080", nil)
}

/*
Registers a book in the database
*/
func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Gets JSON data from request body and puts it in object
	decoder := json.NewDecoder(r.Body)

	// Creates new Book to parse JSON into
	var book Book
	err := decoder.Decode(&book)
	if err != nil {
		fmt.Fprintln(w, "Error decoding JSON data: ", err)
	}
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, book)
}

