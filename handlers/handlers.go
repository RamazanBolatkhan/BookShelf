package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/models"
	"restapi/storage/database"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []models.Book
	if err := database.DB.Find(&books).Error; err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	var book models.Book

	if err := database.DB.First(&book, "id = ?", id).Error; err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)

}

func PostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&book).Error; err != nil {
		http.Error(w, "Failed to create to store in bookshelf", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//Validation of the 'id' in the request
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid ID format", http.StatusBadRequest)
		return
	}

	var book models.Book
	// Find the book from the database with matching id
	if err := database.DB.First(&book, id).Error; err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	var updatedbook models.Book
	//Decodes the updates data from request body
	if err := json.NewDecoder(r.Body).Decode(&updatedbook); err != nil {
		http.Error(w, "Invalid format", http.StatusBadRequest)
		return
	}

	book.Title = updatedbook.Title
	book.Isbn = updatedbook.Isbn
	book.Author = updatedbook.Author
	//Save updated info in the DB
	if err := database.DB.Save(&book).Error; err != nil {
		http.Error(w, "Could not save in the DB", http.StatusInternalServerError)
		return
	}

	//Respond with updated info
	json.NewEncoder(w).Encode(book)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Validation of the 'id' in the request
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	var book models.Book

	//Checking is this book exsist
	if err := database.DB.First(&book, id).Error; err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	//Deleting this book from the database
	if err := database.DB.Delete(&book, id).Error; err != nil {
		http.Error(w, "Could not delete", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted successfully"))
}
