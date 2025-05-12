package handler

import (
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
