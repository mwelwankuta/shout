package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	database "github.com/mwelwankuta/gotut/database"
)

func InitializeServer() {
	r := mux.NewRouter()

	r.HandleFunc("/post", Post).Methods("POST")
	r.HandleFunc("/home", Home).Methods("GET")
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	database.InitializeDatabase()
	InitializeServer()
}
