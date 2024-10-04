package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Post struct {
	ID      string `json:"id"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Image   string `json:"image"`
	Created string `json:"created"`
	UserID  string `json:"user_id"`
}

var posts = make(map[string]Post)

// Initialize Cloudinary
var cld *cloudinary.Cloudinary

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Cloudinary
	cld, err = cloudinary.NewFromParams(
		os.Getenv("CLOUD_NAME"),
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
	)
	if err != nil {
		panic("Failed to initialize Cloudinary: " + err.Error())
	}
}

// HomePage handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html") // Ensure this path is correct
}

// GetPosts handler
func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// CreatePost handles POST requests to create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // Limit image upload size to 10MB
	if err != nil {
		log.Printf("Error parsing form data: %v", err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Extract form fields
	author := r.FormValue("author") // Now this will be set from localStorage correctly
	content := r.FormValue("content")

	// Log to ensure the values are correct
	log.Printf("Author: %s, Content: %s", author, content)

	// Handle image upload via Cloudinary (same as before)
	file, _, err := r.FormFile("image")
	if err != nil {
		log.Printf("Error retrieving the file: %v", err)
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Upload the image to Cloudinary
	uploadResult, err := cld.Upload.Upload(r.Context(), file, uploader.UploadParams{Folder: "posts_images"})
	if err != nil {
		log.Printf("Error uploading image to Cloudinary: %v", err)
		http.Error(w, "Error uploading the image to Cloudinary: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new post with the Cloudinary image URL
	post := Post{
		ID:      uuid.New().String(),
		Content: content,
		Author:  author, // Now correctly set from localStorage
		Image:   uploadResult.SecureURL,
		Created: time.Now().Format(time.RFC3339),
		UserID:  r.FormValue("user_id"), // Make sure to get the user ID if needed
	}

	// Add the post to the in-memory database
	posts[post.ID] = post

	// Respond with the created post
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// UpdatePost handler
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if the post exists
	post, exists := posts[id]
	if !exists {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	// Parse the multipart form data (allow for an image update)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Update the fields
	if author := r.FormValue("author"); author != "" {
		post.Author = author
	}
	if content := r.FormValue("content"); content != "" {
		post.Content = content
	}

	// Handle image upload (optional update of the image)
	file, _, err := r.FormFile("image")
	if err == nil { // If an image is provided
		defer file.Close()

		// Upload the new image to Cloudinary
		uploadResult, err := cld.Upload.Upload(r.Context(), file, uploader.UploadParams{Folder: "posts_images"})
		if err != nil {
			http.Error(w, "Error uploading the image to Cloudinary", http.StatusInternalServerError)
			return
		}

		// Update the image URL in the post
		post.Image = uploadResult.SecureURL
	}

	// Update the post in the in-memory database
	posts[id] = post

	// Respond with the updated post
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// DeletePost handler
func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if the post exists
	if _, exists := posts[id]; exists {
		delete(posts, id)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Post not found", http.StatusNotFound)
	}
}
