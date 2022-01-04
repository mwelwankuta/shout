package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mwelwankuta/gotut/database"
	"github.com/mwelwankuta/gotut/routes"
)

func main() {

	router := mux.NewRouter()

	database.Connect()
	routes.Setup(router)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
