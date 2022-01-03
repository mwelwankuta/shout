package main

import (
	"encoding/json"
	"net/http"

	"github.com/mwelwankuta/gotut/database"
	"github.com/mwelwankuta/gotut/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []models.Post

	database.DB.Find(&posts)
	err := json.NewEncoder(w).Encode(posts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	database.DB.Create(&post)
	json.NewEncoder(w).Encode(post)
}
