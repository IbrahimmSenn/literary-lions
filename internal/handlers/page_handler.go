package handlers

import (
	"log"
	"net/http"
)

// HomeHandler handles the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Oi-oi, something went wrong", http.StatusMethodNotAllowed)
		return
	}

	data := PageData{
		Result: "",
		Status: "200",
	}

	err := Templates.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}

// ErrorHandler handles the error page
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/error" {
		http.NotFound(w, r)
		return
	}

	data := PageData{
		Result: "",
		Status: "200",
	}

	err := Templates.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}
