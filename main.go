package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)

	// Start the web server on port 8080
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
