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

	utils.SendMessage(w, users)
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

	utils.SendMessage(w, user)
}

type CreateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var inputUser CreateUserInput

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&inputUser)
	if err != nil {
		utils.SendError(w, "Invalid user input")
		return
	}

	err = models.CreateUser(inputUser.Username, inputUser.Email, inputUser.Password)
	if err != nil {
		utils.SendError(w, err.Error())
		return
	}

	utils.SendMessage(w, "User created successfully", http.StatusCreated)
}

type SignInUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignInUser(w http.ResponseWriter, r *http.Request) {
	var inputUser SignInUserInput

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&inputUser)
	if err != nil {
		utils.SendError(w, "Invalid user input")
		return
	}

	user, err := models.GetUserByUsername(inputUser.Username)
	if err != nil {
		utils.SendError(w, "Invalid username")
		return
	}

	if !utils.IsPasswordValid(user.Password, inputUser.Password) {
		utils.SendError(w, "Invalid password")
		return
	}

	utils.SendMessage(w, "User signed in successfully")
}
