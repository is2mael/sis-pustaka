package response

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse adalah struktur respons untuk sukses
type SuccessResponse struct {
	Status  string      `json:"status"`           // Selalu "success"
	Message string      `json:"message"`          // Pesan singkat seperti "Berhasil mengambil data"
	Data    interface{} `json:"data,omitempty"`   // Data yang dikembalikan, bisa null
}

// ErrorResponse adalah struktur respons untuk error
type ErrorResponse struct {
	Status  string `json:"status"`  // Selalu "error"
	Message string `json:"message"` // Penjelasan error
}

// WriteSuccess mengirim response sukses dengan format JSON
func WriteSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

// WriteError mengirim response error dengan format JSON
func WriteError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Status:  "error",
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}
