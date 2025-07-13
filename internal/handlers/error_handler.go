package handlers

import (
	"net/http"
)

// NotFoundHandler handles 404 errors for invalid URLs
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html")

	// Use the pre-parsed templates
	err := Templates.ExecuteTemplate(w, "error404.html", nil)
	if err != nil {
		// Fallback if template fails to execute
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}
}

// InternalServerErrorHandler handles 500 errors for server issues
func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/html")

	// Use the pre-parsed templates
	err := Templates.ExecuteTemplate(w, "error500.html", nil)
	if err != nil {
		// Fallback if template fails to execute
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}
}