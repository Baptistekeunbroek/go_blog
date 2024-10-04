package models

// User model
type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// In-memory database for users (replace with actual DB in production)
var Users = make(map[string]User)
