package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/models"
	"restapi/storage/database"
	"strconv"

	"github.com/gorilla/mux"
)

// Get all books
// @Summary Get all Books
// @Description Getting all the books stored in the "Shelf"
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []models.Book
	if err := database.DB.Find(&books).Error; err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(books)
}

// Get book by id
// @Summary Get Book by id
// @Description Getting certain book by ID
// @Tags book
// @Accept json
// @Param  id path int  true  "Book ID"
// @Produce json
// @Success 200 {book} models.Book
// @Router /books/{id} [get]
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Ivalid id format", http.StatusBadRequest)
	}
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)

}

// PostBook godoc
// @Summary Create a new book
// @Description Add a new book to the database
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body models.Book true "Book Data"
// @Success 201 {object} models.Book
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to create book"
// @Router /books [post]
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

// UpdateBook
// @Summary Update Book
// @Description Update a book in the database
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body models.Book true "Book Data"
// @Param id path int true "Book ID"
// @Success 201 {object} models.Book
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to update book"
// @Router /books/{id} [put]
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

// DeleteBook
// @Summary Delete Book
// @Description Delete a book in the database
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 201 {object} models.Book
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to delete book"
// @Router /books/{id} [delete]
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
