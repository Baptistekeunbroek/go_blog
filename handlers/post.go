package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Created string `json:"created"`
}

var posts = make(map[string]Post)

// HomePage handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

// GetPosts handler
func GetPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}

// CreatePost handler
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	posts[newPost.ID] = newPost
	json.NewEncoder(w).Encode(newPost)
}

// UpdatePost handler
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)

	if _, exists := posts[id]; exists {
		posts[id] = updatedPost
		json.NewEncoder(w).Encode(updatedPost)
	} else {
		http.Error(w, "Post not found", http.StatusNotFound)
	}
}

// DeletePost handler
func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, exists := posts[id]; exists {
		delete(posts, id)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Post not found", http.StatusNotFound)
	}
}
