package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	//custom packages
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = map[string]User{}
var jwtKey = []byte("secret")

func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal Server Error : cannot Hash password", http.StatusInternalServerError)
		return
	}
	id := uuid.New().String()
	user := User{ID: id, Email: credentials.Email, Password: string(hash)}
	users[id] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid Input: Check user details", http.StatusBadRequest)
		return
	}
	for _, user := range users {
		if user.Email == credentials.Email && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) == nil {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email": user.Email,
				"exp":   time.Now().Add(time.Hour * 72).Unix(),
			})
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				http.Error(w, "Internal Server Error : cannot Sign Token", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
			return
		}
	}

	http.Error(w, "User not found || Unauthorized", http.StatusNotFound)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, ok := users[vars["id"]]
	if !ok {
		http.Error(w, "User not found || Unauthorized", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", registerUserHandler).Methods("POST")
	r.HandleFunc("login", loginUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	fmt.Println("Listening on port 8080: ")
	log.Fatal(http.ListenAndServe(":8080", r))
}
