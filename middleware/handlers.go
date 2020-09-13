package middleware

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"net/http"
	"github.com/chanum/restapi/models"
	"github.com/gorilla/mux"
)

// Init book var as a slice Book struct
var books []models.Book

func InitMockData() {
	// Mock Data @todo implement a DB
	books = append(books, models.Book{ID:"1", Isbn:"12345", Title:"Moby Dick", Author: &models.Author{Firstname:"Herman", Lastname:"Melville"}})
	books = append(books, models.Book{ID:"2", Isbn:"2222", Title:"The Raven", Author: &models.Author{Firstname:"Edgar Allan", Lastname:"Poe"}})
	books = append(books, models.Book{ID:"3", Isbn:"333334", Title:"Ficciones", Author: &models.Author{Firstname:"Jorge Luis", Lastname:"Borges"}})
}

// Get All Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Book{})
}

// Create a New Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete Book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
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