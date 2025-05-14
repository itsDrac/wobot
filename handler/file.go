package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/itsDrac/wobot/types"
)

// UploadFile godoc
// @Summary Upload a file
// @Description Upload a file to the server
// @Tags Files
// @Accept multipart/form-data
// @Param file formData file true "File to upload"
// @Security Bearer
// @Success 201 {object} types.ApiResponse
// @Failure 409 {object} string "Not enough storage space"
// @Failure 500 {object} string "Internal server error"
// @Router /upload [post]
func (h ChiHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Cap for request size.
	const maxSizePerRequest = 10 << 20 // 10MB
	r.Body = http.MaxBytesReader(w, r.Body, maxSizePerRequest)
	// Storing 100KB in request rest in disk
	if err := r.ParseMultipartForm(100 << 10); err != nil {
		http.Error(w, "Failed to parse file", http.StatusInternalServerError)
		return
	}
	file, fileInfo, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Fail to get from from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	if err := h.Service.File.UploadFile(r.Context(), file, fileInfo); err != nil {
		if strings.Contains(err.Error(), "not enough storage space") {
			http.Error(w, "Not enough storage space", http.StatusBadRequest)
			return
		}
		http.Error(w, "Can not upload file", http.StatusInternalServerError)
		log.Println("Error in uploading file", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := types.ApiResponse{
		Message: "File uploaded successfully",
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
		return
	}

}

// GetStorage godoc
// @Summary Get remaining storage
// @Description Get remaining storage for the user
// @Tags Files
// @Accept json
// @Security Bearer
// @Success 200 {object} types.ApiResponse "Remaining storage: 100.0 MB"
// @Failure 500 {object} string "Internal server error"
// @Router /storage/remaining [get]
func (h ChiHandler) GetStorage(w http.ResponseWriter, r *http.Request) {

	remainingStorage, err := h.Service.File.GetRemainingStorage(r.Context())
	if err != nil {
		http.Error(w, "Failed to get remaining storage", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := types.ApiResponse{
		Message: fmt.Sprintf("Remaining storage: %s", remainingStorage),
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
		return
	}
}

// GetFiles godoc
// @Summary Get files
// @Description Get files for the user
// @Tags Files
// @Accept json
// @Param limit query int false "Limit of files to return"
// @Param offset query int false "Offset for pagination"
// @Security Bearer
// @Success 200 {object} types.Files "List of files"
// @Failure 500 {object} string "Internal server error"
// @Router /files [get]
func (h ChiHandler) GetFiles(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit := query.Get("limit")
	offset := query.Get("offset")
	if limit == "" {
		limit = "10"
	}
	if offset == "" {
		offset = "0"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Invalid limit value", http.StatusBadRequest)
		return
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		http.Error(w, "Invalid offset value", http.StatusBadRequest)
		return
	}
	log.Printf("Limit: %d, Offset: %d", limitInt, offsetInt)
	resp, err := h.Service.File.GetFiles(r.Context(), limitInt, offsetInt)
	if err != nil {
		if strings.Contains(err.Error(), "user not found in context") {
			http.Error(w, "User not found in context", http.StatusBadRequest)
			return
		}
		if strings.Contains(err.Error(), "failed to get files") {
			http.Error(w, "Failed to get files", http.StatusInternalServerError)
			return
		}
		if strings.Contains(err.Error(), "failed to get file info") {
			http.Error(w, "Failed to get file info", http.StatusInternalServerError)
			return
		}
		http.Error(w, "Failed to get files", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
		return
	}
}
