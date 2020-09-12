package main

import(
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

/*****************************************************************************************
 * 
*****************************************************************************************/
// Book Struct (Model)
type Book struct {
	ID		string `json:"id"`
	Isbn	string `json:"isbn"`
	Title	string `json:"title"`
	Author	*Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	Firstname	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`
}

/*****************************************************************************************
 * 
*****************************************************************************************/
// Init book var as a slice Book struct
var books []Book

/*****************************************************************************************
 * 
*****************************************************************************************/
// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

/*****************************************************************************************
 * Main
*****************************************************************************************/
func main ()  {
	// Init Router
	router := mux.NewRouter().StrictSlash(true)

	// Mock Data @todo implement a DB
	books = append(books, Book{ID:"1", Isbn:"12345", Title:"Moby Dick", Author: &Author{Firstname:"Herman", Lastname:"Melville"}})
	books = append(books, Book{ID:"2", Isbn:"2222", Title:"The Raven", Author: &Author{Firstname:"Edgar Allan", Lastname:"Poe"}})
	books = append(books, Book{ID:"3", Isbn:"333334", Title:"Ficciones", Author: &Author{Firstname:"Jorge Luis", Lastname:"Borges"}})

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run the server
	log.Fatal(http.ListenAndServe(":8000", router))
}