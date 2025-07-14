#!/bin/bash

echo "🧪 Testing Literary Lions Forum - 404 Error Handling"
echo "=================================================="

BASE_URL="http://localhost:8080"

# Function to test a URL and check response
test_url() {
    local url="$1"
    local expected_status="$2"
    local description="$3"
    
    echo -n "Testing $description ($url)... "
    
    status=$(curl -s -o /dev/null -w "%{http_code}" "$url")
    
    if [ "$status" -eq "$expected_status" ]; then
        echo "✅ PASS ($status)"
    else
        echo "❌ FAIL (Expected: $expected_status, Got: $status)"
    fi
}

echo ""
echo "📋 Testing Valid Routes (should return 200):"
echo "--------------------------------------------"

test_url "$BASE_URL/" 200 "Home page"
test_url "$BASE_URL/forum" 200 "Forum page"
test_url "$BASE_URL/login" 200 "Login page"
test_url "$BASE_URL/register" 200 "Register page"
test_url "$BASE_URL/newpost" 200 "New post page (might redirect)"
test_url "$BASE_URL/profile" 200 "Profile page (might redirect)"

echo ""
echo "🚫 Testing Invalid Routes (should return 404):"
echo "----------------------------------------------"

test_url "$BASE_URL/nonexistent" 404 "Non-existent page"
test_url "$BASE_URL/random-page" 404 "Random page"
test_url "$BASE_URL/invalid/path" 404 "Invalid path"
test_url "$BASE_URL/forum/invalid" 404 "Invalid forum path"
test_url "$BASE_URL/api/something" 404 "Non-existent API"
test_url "$BASE_URL/post/999999" 404 "Non-existent post"

echo ""
echo "📁 Testing Static Files:"
echo "------------------------"

test_url "$BASE_URL/static/reset.css" 200 "CSS file"
test_url "$BASE_URL/static/nonexistent.css" 404 "Non-existent CSS"
test_url "$BASE_URL/media/logo2.png" 200 "Media file (if exists)"

echo ""
echo "🔍 Testing Edge Cases:"
echo "----------------------"

test_url "$BASE_URL/POST" 404 "Uppercase path"
test_url "$BASE_URL/forum/" 200 "Trailing slash"
test_url "$BASE_URL//forum" 404 "Double slash"

echo ""
echo "🏁 Testing complete!"