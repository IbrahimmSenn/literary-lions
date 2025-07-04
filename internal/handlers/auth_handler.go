package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"literary-lions-forum/internal/database"
	"literary-lions-forum/internal/models"
	"literary-lions-forum/internal/repository/sqlite"
	"literary-lions-forum/internal/utils"

	"github.com/google/uuid"
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

		userRepo := sqlite.NewUserRepository(db)

		// Check if email already exists(requirement)
		existingUser, _ := userRepo.GetUserByEmail(email)
		if existingUser != nil {
			data := PageData{
				Error:  "Email already taken",
				Status: "400",
			}
			Templates.ExecuteTemplate(w, "register.html", data)
			return
		}
		// Check if username already exists(requirement)
		existingUserByUsername, _ := userRepo.GetUserByUsername(username)
		if existingUserByUsername != nil {
			data := PageData{
				Error:  "Oops! This username is already taken. Try another one!",
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
		// Simple debug to see if i can actually register PLEASE REMEMBER TO DELETE IT!!!!!!!
		log.Printf("✅ User successfully registered: ID=%d, Username=%s, Email=%s", newUser.ID, newUser.Username, newUser.Email)

		// Registration successful - redirect to login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Handle unsupported methods
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}

	// Handle GET request - show login form
	if r.Method == http.MethodGet {
		data := PageData{
			Result: "",
			Status: "200",
		}

		err := Templates.ExecuteTemplate(w, "login.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Printf("Template error: %v", err)
		}
		return
	}

	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Get form values
		usernameOrEmail := r.FormValue("email")
		password := r.FormValue("password")

		// Basic validation
		if usernameOrEmail == "" || password == "" {
			data := PageData{
				Error:  "Email and password are required",
				Status: "400",
			}
			Templates.ExecuteTemplate(w, "login.html", data)
			return
		}

		// Connect to database
		db, err := database.ConnectDB()
		if err != nil {
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Create repositories
		userRepo := sqlite.NewUserRepository(db)
		sessionRepo := sqlite.NewSessionRepository(db)

		// Find user by email
		user, err := userRepo.GetUserByEmail(usernameOrEmail)
		if err != nil {
			data := PageData{
				Error:  "Invalid email or password",
				Status: "400",
			}
			Templates.ExecuteTemplate(w, "login.html", data)
			return
		}

		// Check password
		if !utils.CheckPassword(password, user.Password) {
			data := PageData{
				Error:  "Invalid email or password",
				Status: "400",
			}
			Templates.ExecuteTemplate(w, "login.html", data)
			return
		}

		sessionID := generateSessionID()
		session := &models.Session{
			ID:        sessionID,
			UserID:    user.ID,
			ExpiresAt: time.Now().Add(24 * time.Hour),
		}

		// Save session to database
		err = sessionRepo.CreateSession(session)
		if err != nil {
			log.Printf("Error creating session: %v", err)
			http.Error(w, "Login failed", http.StatusInternalServerError)
			return
		}

		// Set session cookie(check later on if there is any other way to do it)
		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Expires:  session.ExpiresAt,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)

		log.Printf("✅ User logged in successfully: %s (ID: %d)", user.Username, user.ID)

		http.Redirect(w, r, "/forum", http.StatusSeeOther)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
func generateSessionID() string {
	return uuid.New().String()
}
