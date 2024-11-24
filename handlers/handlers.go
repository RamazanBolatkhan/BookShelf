package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/models"
	"restapi/storage"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(storage.Books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid ID format", http.StatusBadRequest)
		return
	}

	for _, item := range storage.Books {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Book Not found", http.StatusNotFound)

}

func PostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	book.ID = storage.NextID
	storage.NextID++
	storage.Books = append(storage.Books, book)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid ID format", http.StatusBadRequest)
		return
	}

	for index, item := range storage.Books {
		if item.ID == id {
			var book models.Book
			if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
				http.Error(w, "Invalid Input", http.StatusBadRequest)
				return
			}

			book.ID = id

			storage.Books[index] = book

			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	for index, item := range storage.Books {
		if item.ID == id {
			storage.Books = append(storage.Books[:index], storage.Books[index+1:]...)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusBadRequest)
}
