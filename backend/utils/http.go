package utils

import (
	"log"
	"net/http"
)

func SendError(w http.ResponseWriter, msg string) {
	log.Println("Error:", msg)
	http.Error(w, msg, http.StatusInternalServerError)
}
