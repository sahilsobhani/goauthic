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
        log.Println(" No .env file found, reading from environment")
    }

    database := db.Connect()
    defer database.Close()


	// Initialize database schema
	if err := db.InitializeSchema(database); err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}

	// Initialize handlers
	h := handlers.NewHandler(database)

	// Set up routes
	r := mux.NewRouter()
	r.HandleFunc("/register", h.RegisterUser).Methods("POST")
	r.HandleFunc("/login", h.LoginUser).Methods("POST")
	r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
	
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", r))
}
