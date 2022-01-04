package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mwelwankuta/shout/server/database"
	"github.com/mwelwankuta/shout/server/routes"
)

func main() {
	godotenv.Load("./.env")
	env := os.Getenv("PORT")

	var PORT string = "8080"

	if len(env) > 0 {
		PORT = env
	}

	router := mux.NewRouter()

	err := database.Connect()
	if err != nil {
		panic("Could not connect to database")
	}

	routes.Setup(router)

	log.Println("Server starting on :" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
