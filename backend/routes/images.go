package routes

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"backend/models"
	"backend/utils"
)

type CreateImageInput struct {
	Image string `json:"image"`
}

func CreateImage(w http.ResponseWriter, r *http.Request) {
	var input CreateImageInput

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, "Invalid user ID")
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&input)
	if err != nil {
		utils.SendError(w, "Invalid input")
		return
	}

	imageData, err := base64.StdEncoding.DecodeString(input.Image)
	if err != nil {
		utils.SendError(w, "Invalid image data")
		return
	}

	err = models.CreateImage(id, imageData)
	if err != nil {
		utils.SendError(w, err.Error())
		return
	}

	utils.SendMessage(w, "Image created successfully", http.StatusCreated)
}
