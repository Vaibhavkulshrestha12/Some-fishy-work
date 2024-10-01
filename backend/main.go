package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "time"
    "os"
    "github.com/joho/godotenv"
)

type Resource struct {
    Title       string `json:"title"`
    URL         string `json:"url"`
    Icon        string `json:"icon"`
    Description string `json:"description"`
}

var resources = []Resource{
    {
        Title:       "LeetCode",
        URL:         "https://leetcode.com",
        Icon:        "https://leetcode.com/favicon.ico",
        Description: "A platform for practicing coding problems.",
    },
    {
        Title:       "The Odin Project",
        URL:         "https://www.theodinproject.com",
        Icon:        "https://www.theodinproject.com/favicon.ico",
        Description: "Free full-stack curriculum.",
    },
}

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

var users = map[string]string{}

func enableCORS(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func main() {
    // Load environment variables
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Set up routes
    http.HandleFunc("/", homePage)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./build/static"))))

    http.HandleFunc("/api/resources", getResources)
    http.HandleFunc("/api/register", registerUser)
    http.HandleFunc("/api/login", loginUser)

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }

    w.Write([]byte("Welcome to the Resource Links API"))
}

func getResources(w http.ResponseWriter, r *http.Request) {
    enableCORS(w)
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(resources); err != nil {
        http.Error(w, "Failed to encode resources", http.StatusInternalServerError)
    }
}

func registerUser(w http.ResponseWriter, r *http.Request) {
    enableCORS(w)

    if r.Method == "OPTIONS" {
        return
    }

    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    if _, exists := users[user.Username]; exists {
        http.Error(w, "User already exists", http.StatusConflict)
        return
    }
    users[user.Username] = user.Password
    w.Write([]byte("User registered successfully"))
}

func loginUser(w http.ResponseWriter, r *http.Request) {
    enableCORS(w)

    if r.Method == "OPTIONS" {
        return
    }

    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    if password, exists := users[user.Username]; !exists || password != user.Password {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    // Retrieve the JWT secret from environment variables
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        http.Error(w, "JWT secret is not set", http.StatusInternalServerError)
        return
    }

    tokenString, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
