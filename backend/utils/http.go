package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func SendError(w http.ResponseWriter, msg string, status ...int) {
	statusCode := http.StatusInternalServerError
	if len(status) > 0 {
		statusCode = status[0]
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{Message: msg}
	json.NewEncoder(w).Encode(response)
}

type ValidResponse struct {
	Data any `json:"data"`
}

func SendMessage(w http.ResponseWriter, data any, status ...int) {
	statusCode := http.StatusOK
	if len(status) > 0 {
		statusCode = status[0]
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ValidResponse{Data: data}
	json.NewEncoder(w).Encode(response)
}
