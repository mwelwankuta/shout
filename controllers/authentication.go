package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/mwelwankuta/shout/database"
	"github.com/mwelwankuta/shout/models"
	"github.com/mwelwankuta/shout/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "something wrong",
		})
		return
	}

	var user models.User

	database.DB.First(&user, "email=?", data["email"])

	if user.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "user not found",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "password not correct",
		})
		return
	}

	expiry := time.Now().Add(time.Hour * 24).Unix()

	token, err := utils.GenerateToken(user.Id, expiry)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "invalid token",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
		"user":  user.Email,
	})

}
