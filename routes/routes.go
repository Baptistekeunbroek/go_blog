package routes

import (
	"net/http"

	"go_blog/handlers"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes all the routes for the blog
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Home route (renders the home page)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	// Auth routes (login and register forms)
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/login.html")
	}).Methods("GET")
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/register.html")
	}).Methods("GET")
	r.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/profile.html")
	}).Methods("GET")

	// Auth API routes (for handling form submissions)
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/register", handlers.Register).Methods("POST")

	// Posts API routes
	r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", handlers.DeletePost).Methods("DELETE")

	// Serve static files (CSS, JS)
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return r
}
