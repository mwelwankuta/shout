package routes

import (
	"github.com/gorilla/mux"
	"github.com/mwelwankuta/gotut/controllers"
)

func Setup(r *mux.Router) {

	r.HandleFunc("/api/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login).Methods("POST")

}
