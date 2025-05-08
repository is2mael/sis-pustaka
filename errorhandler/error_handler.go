package errorhandler

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse struct standar untuk response error
type ErrorResponse struct {
	Message string `json:"message"`
}

// writeError mengirim error dengan format JSON
func WriteError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := ErrorResponse{
		Message: message,
	}

	json.NewEncoder(w).Encode(resp)
}
