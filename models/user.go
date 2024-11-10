package models

// User struct represents the schema for the users table
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password,omitempty"` // omitempty to avoid exposing password in responses
    Role     string `json:"role"`               // either "receptionist" or "doctor"
}
