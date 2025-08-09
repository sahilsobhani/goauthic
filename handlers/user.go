package handlers

package handlers

import (
"encoding/json"
"net/http"

"github.com/google/uuid"
"golang.org/x/crypto/bcrypt"

"goauthic/models"
"goauthic/storage"
"goauthic/utils"

"github.com/gorilla/mux"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Cannot hash password", http.StatusInternalServerError)
		return
	}
	id := uuid.New().String()
	user := models.User{ID: id, Email: credentials.Email, Password: string(hash)}
	storage.Users[id] = user

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	for _, user := range storage.Users {
		if user.Email == credentials.Email && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) == nil {
			token, err := utils.GenerateToken(user.Email)
			if err != nil {
				http.Error(w, "Error signing token", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{"token": token})
			return
		}
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, ok := storage.Users[vars["id"]]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

