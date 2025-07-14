package handlers

import (
	"net/http"
)

// GetDarkModePreference checks the user's dark mode preference from cookies
func GetDarkModePreference(r *http.Request) bool {
	cookie, err := r.Cookie("darkMode")
	if err != nil {
		return false // default to light mode if no cookie exists
	}
	return cookie.Value == "true"
}

// SetDarkModeCookie sets the dark mode preference cookie
func SetDarkModeCookie(w http.ResponseWriter, darkMode bool) {
	value := "false"
	if darkMode {
		value = "true"
	}
	
	cookie := &http.Cookie{
		Name:     "darkMode",
		Value:    value,
		Path:     "/",
		MaxAge:   86400 * 365, // Cookie expires in 1 year
		HttpOnly: false,      
		SameSite: http.SameSiteLaxMode,
		Secure:   false,       // Set to true if using HTTPS in production
	}
	http.SetCookie(w, cookie)
}

// ToggleDarkModeHandler handles POST requests to toggle dark mode
func ToggleDarkModeHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Get current dark mode preference
	currentMode := GetDarkModePreference(r)
	
	// Toggle the preference
	newMode := !currentMode
	
	// Set the new cookie
	SetDarkModeCookie(w, newMode)
	
	// Redirect back to the page the user came from
	// If no referer, redirect to home page
	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = "/"
	}
	
	http.Redirect(w, r, referer, http.StatusSeeOther)
}