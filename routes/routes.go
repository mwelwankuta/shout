package routes

import (
	"github.com/gorilla/mux"
	"github.com/mwelwankuta/shout/controllers"
)

func Setup(r *mux.Router) {

	r.HandleFunc("/api/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login).Methods("POST")

	r.HandleFunc("/api/shouts", controllers.Home).Methods("GET")
	r.HandleFunc("/api/delete", controllers.Delete).Methods("POST")
	r.HandleFunc("/api/upload", controllers.ImageUpload).Methods("POST")

	r.HandleFunc("/api/comment", controllers.Comment).Methods("POST")
	r.HandleFunc("/api/like", controllers.Like).Methods("POST")
	r.HandleFunc("/api/delete", controllers.Delete).Methods("DELETE")
}
