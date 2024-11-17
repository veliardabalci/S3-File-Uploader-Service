package main

import (
	"encoding/json"
	"net/http"
)

// FileUploadHandler handles file uploads
func FileUploadHandler(uploader Uploader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20) // 10 MB
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		url, err := uploader.UploadFile(r.Context(), file, handler.Filename)
		if err != nil {
			http.Error(w, "File upload failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := map[string]string{"url": url}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}
}
