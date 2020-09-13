package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/chanum/restapi/router"
)

func main ()  {
	// Init Router
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")

	// Run the server
	log.Fatal(http.ListenAndServe(":8000", r))
}