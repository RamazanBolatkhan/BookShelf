package main

import (
	"log"
	"net/http"
	"restapi/handlers"
	"restapi/models"
	"restapi/storage"

	"github.com/gorilla/mux"
)

func main() {
	storage.Books = append(storage.Books, models.Book{
		ID:    1,
		Isbn:  "438227",
		Title: "Book One",
		Author: &models.Author{
			Firstname: "John",
			Lastname:  "Deer",
		},
	})
	storage.Books = append(storage.Books, models.Book{
		ID:    2,
		Isbn:  "454555",
		Title: "Book two",
		Author: &models.Author{
			Firstname: "Ramazan",
			Lastname:  "Bolatkhan",
		},
	})

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
