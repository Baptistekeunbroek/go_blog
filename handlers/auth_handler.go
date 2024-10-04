package handlers

import (
	"encoding/json"
	"go_blog/models"
	"net/http"

	"github.com/google/uuid"
)

// In-memory storage for users (simulating a database)
var users = make(map[string]models.User) // username -> User

// Register a new user
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if the username already exists
	if _, exists := users[user.Username]; exists {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// Generate a unique UserID
	user.UserID = uuid.New().String() // Assign a unique ID

	// Store the user in the in-memory map
	users[user.Username] = user

	// Respond with the success message and the user ID and username
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "User registered successfully",
		"user_id":  user.UserID,
		"username": user.Username,
	})
}

// Login an existing user
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if the user exists
	storedUser, exists := users[user.Username]
	if !exists || storedUser.Password != user.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Respond with success and include the UserID and Username
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Login successful",
		"user_id":  storedUser.UserID,
		"username": storedUser.Username,
	})
}
