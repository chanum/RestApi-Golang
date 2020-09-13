package router

import (
	"github.com/chanum/restapi/middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router { 
	// Init Router
	router := mux.NewRouter().StrictSlash(true)

	// TODO: replace with a DB
	middleware.InitMockData()
	
	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", middleware.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", middleware.GetBook).Methods("GET")
	router.HandleFunc("/api/books", middleware.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", middleware.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", middleware.DeleteBook).Methods("DELETE")

	return router
}