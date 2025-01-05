package routes

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"

	"backend/models"
	"backend/utils"
)

type CreateImageInput struct {
	Image   string `json:"image"`
	Overlay string `json:"overlay"`
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

	userIDStr := r.Header.Get("X-User-ID")
	if userIDStr != idStr {
		utils.SendError(w, "Invalid user ID", http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&input)
	if err != nil {
		utils.SendError(w, "Invalid input")
		return
	}

	backgroundData, err := base64.StdEncoding.DecodeString(input.Image)
	if err != nil {
		utils.SendError(w, "Invalid image")
		return
	}

	background, err := png.Decode(bytes.NewReader(backgroundData))
	if err != nil {
		utils.SendError(w, "Invalid image")
		return
	}

	overlayFilepath := fmt.Sprintf("/app/images/%s", input.Overlay)
	overlayFile, err := os.Open(overlayFilepath)
	if err != nil {
		utils.SendError(w, "Invalid overlay image")
		return
	}
	defer overlayFile.Close()

	overlay, err := png.Decode(overlayFile)
	if err != nil {
		utils.SendError(w, "Invalid overlay image")
		return
	}
	overlayResized := resize.Resize(
		uint(background.Bounds().Dx()),
		uint(background.Bounds().Dy()),
		overlay,
		resize.Lanczos3,
	)

	finalImage := image.NewRGBA(background.Bounds())

	draw.Draw(finalImage, background.Bounds(), background, image.Point{}, draw.Over)

	draw.Draw(finalImage, overlayResized.Bounds(), overlayResized, image.Point{}, draw.Over)

	var buf bytes.Buffer
	err = jpeg.Encode(&buf, finalImage, nil)
	if err != nil {
		utils.SendError(w, "Failed to generate final image")
		return
	}

	finalImageData := buf.Bytes()
	err = models.CreateImage(id, finalImageData)
	if err != nil {
		utils.SendError(w, err.Error())
		return
	}

	utils.SendMessage(w, "Image created successfully", http.StatusCreated)
}

func GetImages(w http.ResponseWriter, r *http.Request) {
	images, err := models.GetImages()
	if err != nil {
		utils.SendError(w, err.Error())
		return
	}

	utils.SendMessage(w, images)
}

func GetUserImages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, "Invalid user ID")
		return
	}

	images, err := models.GetUserImages(id)
	if err != nil {
		utils.SendError(w, err.Error())
		return
	}

	utils.SendMessage(w, images)
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, "Invalid image ID")
		return
	}

	imageData, err := models.GetImage(id)
	if err != nil {
		utils.SendError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(imageData)
}
