package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"backend/models"
	"backend/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers()
	if err != nil {
		utils.SendError(w, "Invalid users query")
		return
	}

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, "Invalid user ID")
		return
	}

	user, err := models.GetUser(id)
	if err != nil {
		utils.SendError(w, "Invalid user query")
		return
	}

	json.NewEncoder(w).Encode(user)
}

type NewUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUserInput NewUserInput

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUserInput)
	if err != nil {
		utils.SendError(w, "Invalid user input")
		return
	}

	newUser, err := models.CreateUser(newUserInput.Username, newUserInput.Email, newUserInput.Password)
	if err != nil {
		utils.SendError(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(newUser)
}
