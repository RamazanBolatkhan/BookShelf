package main

import (
	"log"
	"net/http"
	"restapi/handlers"
	"restapi/storage/database"

	"github.com/gorilla/mux"
)

func main() {
	//connecting to database
	database.Connectdb()

	//Init router
	r := mux.NewRouter()

	//register routes {Get; Get by ID; Post; Put; Delete}
	r.HandleFunc("/books", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBookByID).Methods("GET")
	r.HandleFunc("/books", handlers.PostBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
