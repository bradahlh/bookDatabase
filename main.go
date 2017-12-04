package main

import (
	"net/http"
	"os"
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

	// Web server (gets port number from environment variable)
	port := os.Getenv("PORT")
	http.HandleFunc("/register", bookDB.registerHandler)
	http.ListenAndServe(":"+port, nil)
}
