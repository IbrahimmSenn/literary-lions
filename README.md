# Cars-Viewer – OnlyCars
OnlyCars is an intuitive platform where car enthusiasts can explore different car models, search for vehicles based on their preferences and compare specifications side-by-side.

## 1. Summary
### Current situation:

Car enthusiasts often struggle to find comprehensive, organized information about different vehicle models in one place. Existing automotive websites may be cluttered, difficult to navigate or lack proper comparison tools. Users need a streamlined way to:

- Browse different car models and specifications
- Search for vehicles that match their interests
- Compare multiple cars side-by-side
- Access detailed information in an organized format

The challenge is to create a platform that presents complex automotive data in an intuitive, visually appealing manner while maintaining ease of use across all devices.

### Solution:
OnlyCars addresses these needs by providing a dedicated web platform that showcases comprehensive information about different car models, their specifications, manufacturers and more in a user-friendly way. The solution includes:

- Clean, modern interface with card-based design for easy browsing
- Advanced search functionality for quick vehicle discovery
- Detailed car pages with comprehensive specifications
- Side-by-side comparison tool for multiple vehicles
- Filtering by type, manufacturer, country, and year
- Responsive design optimized for all devices

Information is displayed through visually appealing car cards and detailed specification pages, making it easy for users to find and compare vehicles that match their interests.

## 2. Code Structure
```
cars/
├── main.go                    # Application entry point
├── go.mod                     # Go module dependencies
├── README.md                  # Project documentation
├── api/                       # API server (separate service)
├── handlers/                  # HTTP request handlers
│   ├── cardetailshandler.go  # Car details page handler
│   ├── comparehandler.go     # Car comparison handler
│   ├── mainpagehandler.go    # Home page handler
│   ├── searchcarshandle.go   # Search functionality handler
│   └── utils.go              # Handler utilities
├── images/                    # Image assets
├── jsondata/                  # JSON data files
├── static/                    # Static CSS files
│   ├── style.css             # Global styles
│   ├── cardetails.css        # Car details page styles
├── templates/                 # HTML templates
│   ├── carcards.html         # Car card component
│   ├── cardetails.html       # Car details view
│   ├── compare.html          # Car comparison page
│   ├── footer.html           # Site footer component
│   ├── header.html           # Site header component
│   ├── home.html             # Home page template
│   └── search.html           # Search results page
├── utils/                     # Utility functions
│   ├── cookie_helper.go      # Cookie management
│   ├── fetch_data_by_id.go   # Data fetching by ID
│   └── fetch_data.go         # General data fetching
```

## 3. Installation
### Quick Start

Clone the repository
```
bash
git clone <https://gitea.koodsisu.fi/helisavila/cars>
cd cars
```
Start the API server
```
cd api
make build  # First time only
export PORT=3000  # Linux/Mac
# OR
$env:PORT=3000  # Windows PowerShell
make run &
```
The API server will run on http://localhost:3000

Start the web application
```
cd ..  # Back to project root (cars/)
go run .
```
The web application will start on http://localhost:8080

### Access the application
Open your browser and navigate to http://localhost:8080
The API serves data to the web application from port 3000

#### Development Notes
- Leave the API server running in the background
- The web application fetches data from the API server
- Static files are served from the /static directory
- Templates are rendered server-side using Go's html/template package

## 4. Features
- Interactive Home Page with slideshow and car grid.
- Responsive Design for mobile and desktop.
- Smooth hover animations and transitions
- Grid-based responsive layouts
- Clean typography with Lora font family
- Dark theme with blue accent colors
- Intuitive navigation and user flow


## 5. Bonus features (for review)
Filter feature – user can filter cars by vehicle type, make, country and year.
Comparison – user can pick cars to compare on a separate page.
Personalised recommendations – user can see similar cars to the one they are reading details about.
Search tab – user can separately search for cars using keywords.

## 6. Usage

### Browsing Cars
- Visit the home page to see featured cars
- Use filter dropdowns to narrow by type, make, country or year
- Click on any car card to view detailed specifications

### Searching
- Use the search bar on any page
- Enter keywords like car names, manufacturers, or features

### Comparing Cars
- Select "Add to comparison" on car cards
- Choose multiple vehicles
- Click "Compare selected" to see side-by-side specs


## 7. Credits & Attributions
Logo / Car
(https://www.flaticon.com/free-icons/eco-car)
Eco car icons created by Mike Zuidgeest - Flaticon

Search bar icon
(https://www.flaticon.com/free-icons/search-bar)
Search bar icons created by Bharat Icons - Flaticon

Home page photos:
1. (https://unsplash.com/photos/brown-boat-on-snow-covered-ground-near-green-trees-and-snow-covered-mountains-during-daytime-Vc_goI6kwT0)
Alain Gehri - Unsplash 


2. (https://unsplash.com/photos/a-car-is-parked-in-a-field-with-mountains-in-the-background-veikogWYk14)
Spenser Sembrat - Unsplash


3. (https://unsplash.com/photos/a-car-driving-down-a-road-with-a-mountain-in-the-background-s6Xi05cto-o)
Erwi - Unsplash

Fonts
Main font: (https://fonts.google.com/specimen/Lora/license)
Lora - Copyright 2011 The Lora Project Authors
Extra font: (https://fonts.google.com/specimen/Playwrite+DK+Loopet?query=Playwrite)
Playwright - Copyright 2023 The Playwrite Project Authors


## Authors
- [Heli Savila](https://gitea.kood.tech/helisavila) Frontend Development, UI/UX Design
- [Ibrahim Sen](https://gitea.kood.tech/ibrahimsen) Backend Development, Architecture

