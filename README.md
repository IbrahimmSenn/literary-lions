# Book Forum – Bookies
Bookies is an intuitive forum platform where book enthusiasts can discuss literature, share insights and engage in meaningful conversations with fellow readers.

## 1. Summary
### Current situation:

The "Literary Lions" book club was drowning in paper-based discussions that created chaos and confusion. Sticky notes overflowed pages, comments got lost in scribbles and finding past insights felt like navigating a literary labyrinth. New members felt intimidated by the chaotic paper trails, hindering their engagement and knowledge sharing.
The challenge was to design a digital forum that tames the paper pandemonium and fosters a thriving online book club community.

### Solution:
Literary Lions Forum addresses these needs by providing a comprehensive digital platform that transforms chaotic paper discussions into organized, accessible conversations. The solution includes:

- User authentication system with secure registration and login
- Comprehensive post and comment system for rich discussions
- Category-based organization for different books, genres, and topics
- Like/dislike functionality to highlight valuable contributions
- Advanced filtering capabilities for easy content discovery
- Responsive design optimized for all devices
- SQLite database for reliable data management
- Dockerized deployment for easy setup and scalability

## 2. Code Structure
```
literary-lions/
├── cmd/
│   └── main.go                   # Application entry point
├── internal/
│   ├── data/                     # Data initialization
│   │   ├── books.go             # Book data setup
│   │   └── quotes.go            # Quote data setup
│   ├── database/
│   │   └── database.go          # Database connection and setup
│   └── handlers/                # HTTP request handlers
│       ├── auth_handler.go      # Authentication handlers
│       ├── comment_handler.go   # Comment management
│       ├── error_handler.go     # Error page handlers
│       ├── page_handler.go      # Main page handlers
│       ├── post_handler.go      # Post management
│       ├── user_handler.go      # User profile handlers
│       └── vote_handler.go      # Like/dislike functionality
├── models/                      # Data models
│   ├── books.go                # Book model
│   ├── category.go             # Category model
│   ├── comment.go              # Comment model
│   ├── like.go                 # Like/dislike model
│   ├── post.go                 # Post model
│   ├── quote.go                # Quote model
│   ├── session.go              # Session model
│   └── user.go                 # User model
├── repository/                  # Data access layer
│   ├── sqlite/                 # SQLite implementations
│   │   ├── books_repository.go
│   │   ├── category_repository.go
│   │   ├── comment_repository.go
│   │   ├── like_repository.go
│   │   ├── post_repository.go
│   │   ├── quote_repository.go
│   │   ├── session_repository.go
│   │   └── user_repository.go
│   ├── books_repository.go
│   ├── category_repository.go
│   ├── comment_repository.go
│   ├── like_repository.go
│   ├── post_repository.go
│   ├── quote_repository.go
│   ├── session_repository.go
│   └── user_repository.go
├── web/    
│   └── templates/                # HTML templates
│   │   ├── error404.html           # 404 error page
│   │   ├── error500.html           # 500 error page
│   │   ├── footer.html             # Site footer component
│   │   ├── forum.html              # Main forum page
│   │   ├── header.html             # Guest user header
│   │   ├── home.html               # Landing page
│   │   ├── login.html              # User login page
│   │   ├── postview.html           # Individual post view
│   │   ├── register.html           # User registration page
│   │   ├── user_header.html        # Logged-in user header
│   │   ├── user_newpost.html       # New post creation
│   │   └── user_profile.html       # User profile page
│   └── static/                    # CSS stylesheets
│       ├── error.css           # Error page styles
│       ├── footer.css          # Footer styles
│       ├── forum.css           # Forum page styles
│       ├── header.css          # Header styles
│       ├── home.css            # Home page styles
│       ├── login.css           # Login page styles
│       ├── mainstyle.css       # Global styles
│       ├── newpost.css         # New post form styles
│       ├── postview.css        # Post view styles
│       ├── profile.css         # Profile page styles
│       ├── register.css        # Registration page styles
│       └── reset.css           # CSS reset
├── media/                    
│   │   ├── books/ 
│   └── images...               # Logo, lion, error pages etc
├── utils/                      # Utility functions
│   ├── password.go             # Password encryption utilities
│   └── session.go              # Session management utilities
├── database_erd.md             # Entity Relationship Diagram
├── Dockerfile                  # Docker configuration
├── forum.db                    # SQLite database file
├── go.mod                      # Go module dependencies
├── go.sum                      # Go module checksums
└── README.md                   # Project documentation
```

## 3. Installation
### Quick Start

Clone the repository
```
bash
git clone <https://gitea.kood.tech/helisavila/literary-lions>
cd literary-lions
```

### Build and run with Docker (Recommended) 
If you don't have Docker installed, download and install it from:
- [Windows: Docker Desktop for Windows](https://desktop.docker.com/win/main/amd64/Docker%20Desktop%20Installer.exe)
- [Mac: Docker Desktop for Mac](https://desktop.docker.com/mac/main/amd64/Docker.dmg)
- [Linux: Docker Engine](https://docs.docker.com/engine/install/)

```
docker build -t literary-lions .  
docker run -v $(pwd):/host -p 8080:8080 literary-lions
```
### Access the application
Open your browser and navigate to http://localhost:8080

## 4. Features
- User Authentication
- Post Management
- Comment System
- Like/Dislike System
- Category Organization and filtering posts
- Error Handling

## 5. Bonus features (for review)
- User Profiles: personal dashboards showing user activity
- Search tab
- Sorting feature: sort posts by recent, most liked, most disliked
- Responsive Design: optimized for desktop and mobile devices

## 6. Technical Stack
- Backend: Go (Golang) with net/http
- Database: SQLite with go-sqlite3 driver
- Frontend: HTML templates with CSS styling
- Containerization: Docker for deployment
- Session Management: Cookie-based with UUID tokens
- Security: bcrypt password hashing

## 7. Credits & Attributions
#### Fonts
- Logo font: [Barriecito - Copyright 2018 The Barriecito Project Authors](https://fonts.google.com/specimen/Barriecito)
- Main font: [Hepta Slab - Copyright 2018 The Hepta Slab Project Authors](https://fonts.google.com/specimen/Hepta+Slab)

#### Search bar icon
- [Bharat Icons - Flaticon](https://www.flaticon.com/free-icons/search-bar)

#### Photos and images used:
- [Lion - Unsplash](https://www.flaticon.com/free-icon/lion_490125) 
- [Plant - Unsplash](https://www.flaticon.com/free-icon/succulent_3968246) 
- [Plant2 - Unsplash](https://www.flaticon.com/free-icon/alocasia_6670600)
- [Books - Unsplash](https://www.flaticon.com/free-icon/book-stack_3389066)
- [Blobs – Muhammad Abdullah – Freepik](https://www.freepik.com/free-vector/colorful-abstract-fluid-handdrawn-organic-creative-blob-shapes-element_381357218.htm#fromView=keyword&page=1&position=19&uuid=6d9b47bf-4bde-4289-a23c-c517641187c5&query=Blob) 

## Authors
- [Heli Savila](https://gitea.kood.tech/helisavila) Frontend Development, UI/UX Design
- [Ibrahim Sen](https://gitea.kood.tech/ibrahimsen) Backend Development, Architecture, Docker Implementation

