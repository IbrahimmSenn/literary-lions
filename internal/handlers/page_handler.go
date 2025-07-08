package handlers

import (
	"literary-lions-forum/internal/utils"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Oi-oi, something went wrong", http.StatusMethodNotAllowed)
		return
	}

	// Get current user
	currentUser, _ := utils.GetCurrentUser(r)

	data := PageData{
		Result:      "",
		Status:      "200",
		CurrentUser: currentUser,
	}

	// Choose template based on login status
	if currentUser != nil {
		// User is logged in - redirect to profile or show user home
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	} else {
		// User not logged in - show public home
		err := Templates.ExecuteTemplate(w, "home.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Printf("Template error: %v", err)
		}
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
