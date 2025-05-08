package book_handler

import (
	"encoding/json"
	"net/http"
	"sis-pustaka/model"
	"sis-pustaka/storage"
	"sis-pustaka/handler/response"
	"github.com/go-chi/chi/v5"
)

// Get all books
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := storage.GetBookStorage().GetAll()
	response.WriteSuccess(w, http.StatusOK, "Berhasil mengambil semua buku", books)
}

// Get book by ID
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "ID tidak boleh kosong")
		return
	}
	book, found := storage.GetBookStorage().GetByID(id)
	if !found {
		response.WriteError(w, http.StatusNotFound, "Buku tidak ditemukan")
		return
	}
	response.WriteSuccess(w, http.StatusOK, "Buku ditemukan", book)
}

// Create new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		response.WriteError(w, http.StatusBadRequest, "JSON tidak valid")
		return
	}

	// Validasi field tidak boleh kosong
	valid, msg := validateBook(book)
	if !valid {
		response.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	storage.GetBookStorage().Create(book)
	response.WriteSuccess(w, http.StatusCreated, "Buku berhasil dibuat", book)
}

// Update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "ID tidak boleh kosong")
		return
	}

	var updated model.Book
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		response.WriteError(w, http.StatusBadRequest, "JSON tidak valid")
		return
	}

	// Validasi field tidak boleh kosong
	valid, msg := validateBook(updated)
	if !valid {
		response.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	ok := storage.GetBookStorage().Update(id, updated)
	if !ok {
		response.WriteError(w, http.StatusNotFound, "Buku tidak ditemukan")
		return
	}

	response.WriteSuccess(w, http.StatusOK, "Buku berhasil diperbarui", updated)
}

// Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "ID tidak boleh kosong")
		return
	}

	ok := storage.GetBookStorage().Delete(id)
	if !ok {
		response.WriteError(w, http.StatusNotFound, "Buku tidak ditemukan")
		return
	}

	response.WriteSuccess(w, http.StatusOK, "Buku " + id + " berhasil dihapus ", nil)
}
