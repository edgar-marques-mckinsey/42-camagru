package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"backend/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsRequestAuthenticated(r) {
			utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func InitRoutes(router *mux.Router) {
	routerWithAuth := router.NewRoute().Subrouter()
	routerWithAuth.Use(AuthMiddleware)

	routerWithAuth.HandleFunc("/users", GetUsers).Methods("GET")
	routerWithAuth.HandleFunc("/users/{id}", GetUser).Methods("GET")
	routerWithAuth.HandleFunc("/users/{id}", EditUser).Methods("PATCH")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/sign-in", SignInUser).Methods("POST")
	router.HandleFunc("/users/auth", AuthUser).Methods("POST")
}
