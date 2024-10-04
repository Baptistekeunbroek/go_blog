package main

import (
	"fmt"
	"go_blog/handlers" // Ensure this matches your directory structure
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Home page - you can create this handler or remove it if not needed
	r.HandleFunc("/", handlers.HomePage).Methods("GET")

	// Auth routes
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/register", handlers.Register).Methods("POST")

	// Posts API routes
	r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", handlers.DeletePost).Methods("DELETE")

	// Serve static files (HTML, CSS, JS)
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Serve HTML files directly
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/register.html")
	}).Methods("GET")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/login.html")
	}).Methods("GET")

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
