package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ImageUpload(w http.ResponseWriter, r *http.Request) {
	// upload images
	maxMemory := 1024 * 1024 * 10
	r.ParseMultipartForm(int64(maxMemory))

	file, headers, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// TODO: file naming problem
	tempFile, err := ioutil.TempFile("uploads", headers.Filename)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tempFile.Write(fileBytes)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "uploaded file",
	})

}
