package main

import (
	"log"
	"net/http"
	"restapi/handlers"
	"restapi/storage/database"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "restapi/docs"
)

// @title Book Shelf API
// @version 1.0
// @description This is a Pet project to learn Golang, used Postgresql
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://www.linkedin.com/in/ramazan-bolatkhan-a852321b8/
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
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

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
