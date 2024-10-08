package main

import (
	"fmt"
	"go_blog/routes" // Ensure this matches your directory structure
	"log"
	"net/http"
)

func main() {
	// Initialize routes
	r := routes.SetupRoutes()

	// Serve HTML templates and static files (CSS, JS)
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Serve HTML files for register and login pages
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/register.html")
	}).Methods("GET")
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/login.html")
	}).Methods("GET")
	r.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/profile.html")
	}).Methods("GET")

	// Start the server
	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
