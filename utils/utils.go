package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// RespondWithJSON writes JSON response with the given data and HTTP status code
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
    response, err := json.Marshal(payload)
    if err != nil {
        RespondWithError(w, http.StatusInternalServerError, "Error processing JSON response")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(response)
}

// RespondWithError writes JSON error response with the given message and HTTP status code
func RespondWithError(w http.ResponseWriter, status int, message string) {
    RespondWithJSON(w, status, map[string]string{"error": message})
}

// HashPassword hashes a plaintext password (use bcrypt or another secure hashing library)
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Println("Error hashing password:", err)
        return "", err
    }
    return string(hashedPassword), nil
}

// CheckPassword checks if the hashed password matches the plaintext password
func CheckPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}
