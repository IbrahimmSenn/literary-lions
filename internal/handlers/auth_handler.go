package handlers

import (
	"html/template"
	"log"
	"net/http"

	"literary-lions-forum/internal/database"
	"literary-lions-forum/internal/models"
	"literary-lions-forum/internal/repository/sqlite"
	"literary-lions-forum/internal/utils"
)

type PageData struct {
	Input  string
	Result string
	Status string
	Error  string
	Posts  []*models.Post
}

// Templates variable (will be set from main)
var Templates *template.Template

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodGet {
		data := PageData{
			Result: "",
			Status: "200",
		}

		err := Templates.ExecuteTemplate(w, "register.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Printf("Template error: %v", err)
		}
		return
	}

	// Handle POST request - process registration
	if r.Method == http.MethodPost {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Get form values
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if username == "" || email == "" || password == "" {
			data := PageData{
				Error:  "All fields are required",
				Status: "400",
			}
			Templates.ExecuteTemplate(w, "register.html", data)
			return
		}

		// Connect to database
		db, err := database.ConnectDB()
		if err != nil {
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Create user repository
		userRepo := sqlite.NewUserRepository(db)

		// Check if email already exists
		existingUser, _ := userRepo.GetUserByEmail(email)
		if existingUser != nil {
			data := PageData{
				Error:  "Email already registered",
				Status: "400",
			}
			Templates.ExecuteTemplate(w, "register.html", data)
			return
		}

		// Hash the password
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			http.Error(w, "Error processing password", http.StatusInternalServerError)
			return
		}

		//  new user
		newUser := &models.User{
			Username: username,
			Email:    email,
			Password: hashedPassword,
		}

		// Save user to database
		err = userRepo.CreateUser(newUser)
		if err != nil {
			log.Printf("Error creating user: %v", err)
			data := PageData{
				Error:  "Registration failed. Username or email might already exist.",
				Status: "500",
			}
			Templates.ExecuteTemplate(w, "register.html", data)
			return
		}

		// Registration successful - redirect to login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Handle unsupported methods
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// (placeholder for now)
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}

	data := PageData{
		Result: "",
		Status: "200",
	}

	err := Templates.ExecuteTemplate(w, "login.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}
