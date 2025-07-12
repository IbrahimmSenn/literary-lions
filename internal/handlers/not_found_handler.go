package handlers

import (
	"net/http"
)

// NotFoundHandler handles 404 errors for invalid URLs
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>404 - Page Not Found | Bookies Forum</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body { 
            font-family: 'Georgia', serif;
            background: linear-gradient(135deg, #f8f6f0 0%, #ede7d9 100%);
            min-height: 100vh;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            color: #5d4037;
            line-height: 1.6;
        }
        
        .error-container {
            text-align: center;
            max-width: 600px;
            padding: 40px 20px;
        }
        
        .error-image {
            max-width: 100%;
            height: auto;
            margin-bottom: 40px;
            filter: drop-shadow(0 8px 16px rgba(139, 69, 19, 0.15));
        }
        
        .error-code {
            font-size: 4rem;
            font-weight: bold;
            color: #8B4513;
            margin-bottom: 10px;
            text-shadow: 2px 2px 4px rgba(139, 69, 19, 0.1);
        }
        
        .error-title {
            font-size: 2rem;
            color: #8B4513;
            margin-bottom: 20px;
            font-weight: 600;
        }
        
        .error-description {
            font-size: 1.1rem;
            color: #6B4423;
            margin-bottom: 40px;
            line-height: 1.6;
            max-width: 480px;
            margin-left: auto;
            margin-right: auto;
        }
        
        .button-group {
            display: flex;
            gap: 20px;
            justify-content: center;
            flex-wrap: wrap;
            margin-top: 30px;
        }
        
        .btn {
            display: inline-flex;
            align-items: center;
            gap: 8px;
            padding: 12px 24px;
            text-decoration: none;
            border-radius: 25px;
            font-weight: 600;
            transition: all 0.3s ease;
            font-size: 1rem;
            font-family: 'Georgia', serif;
            border: 2px solid transparent;
        }
        
        .btn-primary {
            background: linear-gradient(135deg, #FF8C42, #FF6B35);
            color: white;
            box-shadow: 0 4px 15px rgba(255, 140, 66, 0.3);
        }
        
        .btn-primary:hover {
            transform: translateY(-2px);
            box-shadow: 0 6px 20px rgba(255, 140, 66, 0.4);
        }
        
        .btn-secondary {
            background: transparent;
            color: #8B4513;
            border: 2px solid #FF8C42;
        }
        
        .btn-secondary:hover {
            background: #FF8C42;
            color: white;
            transform: translateY(-2px);
        }
        
        /* Responsive design */
        @media (max-width: 768px) {
            .error-code { 
                font-size: 3rem; 
            }
            .error-title { 
                font-size: 1.5rem; 
            }
            .error-description {
                font-size: 1rem;
                padding: 0 10px;
            }
            .button-group { 
                flex-direction: column; 
                align-items: center; 
            }
            .btn { 
                width: 200px; 
                justify-content: center; 
            }
            .error-image {
                max-width: 90%;
            }
        }
        
        @media (max-width: 480px) {
            .error-container {
                padding: 20px 10px;
            }
            .error-code {
                font-size: 2.5rem;
            }
        }
    </style>
</head>
<body>
    <div class="error-container">
        <img src="/media/error.png" alt="Books scattered around" class="error-image">
        
        <div class="error-code">404</div>
        <h1 class="error-title">Not Found</h1>
        <p class="error-description">
            Please check the URL in the address bar and try again.<br>
            The page you're looking for might have been moved or doesn't exist.
        </p>
        
        <div class="button-group">
            <a href="/" class="btn btn-primary">
                🏠 Home
            </a>
            <a href="/forum" class="btn btn-secondary">
                📚 Forum
            </a>
        </div>
    </div>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(tmpl))
}
