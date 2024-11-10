package controllers

import (
    "encoding/json"
    "myapp/config"   // Import config to access DB connection
    "myapp/services" // Import for Authenticate and JWT generation
    "myapp/models"   // Import for User struct
    "net/http"
    "log"
)

// Register handles the user registration request
func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User

    // Decode the incoming request body to get user details
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Insert the new user into the database
    _, err = config.DB.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)", user.Username, user.Password, user.Role)
    if err != nil {
        log.Printf("Database error: %v", err)
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }
	

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("User registered successfully"))
}

// Login handles the login request
func Login(w http.ResponseWriter, r *http.Request) {
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    // Decode the incoming request body
    err := json.NewDecoder(r.Body).Decode(&credentials)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Authenticate the user and retrieve the user and token
    user, token, err := services.Authenticate(credentials.Username, credentials.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    // Send the JWT token and user role as a response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "token": token,
        "role":  user.Role,  // Send role dynamically
    })
}
