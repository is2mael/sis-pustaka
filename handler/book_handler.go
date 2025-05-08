package book_handler

import (
	"encoding/json" //karna kita mau mengencode data ke json
	"net/http" //untuk menangani request dan response HTTP
	"sis-pustaka/model"	//ambil model yang kita buat sebelumnya	
	"sis-pustaka/storage" //ambil storage yang kita buat sebelumnya
	
	"github.com/go-chi/chi/v5" //import chi untuk routing
)

// Get all books
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := storage.GetBookStorage().GetAll()
	writeJSON(w, http.StatusOK, books)
}

// Get book by ID
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, found := storage.GetBookStorage().GetByID(id)
	if !found {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, book)
}

// Create new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	storage.GetBookStorage().Create(book)
	writeJSON(w, http.StatusCreated, book)
}

// Update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var updated model.Book
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	ok := storage.GetBookStorage().Update(id, updated)
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

// Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	ok := storage.GetBookStorage().Delete(id)
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Deleted"})
}

// Helper: Write JSON response
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}