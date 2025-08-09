package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"goauthic/db"
	"goauthic/handlers"
)

func main() {

	err := godotenv.Load()
    if err != nil {
        log.Println("⚠️ No .env file found, reading from environment")
    }

    database := db.Connect()
    defer database.Close()


	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", r))
}
