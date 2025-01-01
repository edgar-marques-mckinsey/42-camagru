package routes

import (
	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/sign-in", SignInUser).Methods("POST")
	router.HandleFunc("/users/auth", AuthUser).Methods("POST")
}
