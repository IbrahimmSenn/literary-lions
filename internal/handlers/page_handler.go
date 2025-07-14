package handlers

import (
	"literary-lions-forum/internal/models"
	"literary-lions-forum/internal/repository/sqlite"
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
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get current user
	currentUser, _ := utils.GetCurrentUser(r)

	// Get today's quote (keep this working)
	quoteRepo := sqlite.NewQuoteRepository()
	todaysQuote, err := quoteRepo.GetTodaysLiteratureQuote()
	if err != nil {
		log.Printf("Error getting today's quote: %v", err)
		todaysQuote = &models.Quote{
			Text:      "A reader lives a thousand lives before he dies. The man who never reads lives only one.",
			Author:    "George R.R. Martin",
			BookTitle: "A Dance with Dragons",
		}
	}

	// Simple fixed book recommendation - WITH IMAGE
	todaysBook := &models.Book{
		Title:       "The Seven Husbands of Evelyn Hugo",
		Author:      "Taylor Jenkins Reid",
		Genre:       "Contemporary Fiction",
		ImageURL:    "/media/books/evelyn-hugo.png", // Use actual image from media folder
		Description: "A reclusive Hollywood icon reveals her life's secrets to an unknown journalist in this captivating novel.",
	}

	// Prepare data for template
	data := PageData{
		Result:      "",
		Status:      "200",
		CurrentUser: currentUser,
		Quote:       todaysQuote,
		Book:        todaysBook,
		DarkMode:    GetDarkModePreference(r),
	}

	// Render the home template
	err = Templates.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/error" {
		http.NotFound(w, r)
		return
	}

	data := PageData{
		Result: "",
		Status: "200",
		DarkMode: GetDarkModePreference(r),
	}

	err := Templates.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}
