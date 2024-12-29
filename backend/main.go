package main

import (
	"log"
	"net/http"

	"backend/utils"

	"github.com/gorilla/mux"
)

func main() {
	utils.InitDB()
	router := mux.NewRouter()

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
