package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

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
	w.Write([]byte("File uploaded successfully"))

}

func (h ChiHandler) GetStorage(w http.ResponseWriter, r *http.Request) {

	remainingStorage, err := h.Service.File.GetRemainingStorage(r.Context())
	if err != nil {
		http.Error(w, "Failed to get remaining storage", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := struct {
		RemainingStorage string `json:"remaining_storage"`
	}{
		RemainingStorage: remainingStorage,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
		return
	}
}

func (h ChiHandler) GetFiles(w http.ResponseWriter, r *http.Request) {
	resp, err := h.Service.File.GetFiles(r.Context())
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
