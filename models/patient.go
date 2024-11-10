package models

// Patient struct represents the schema for the patients table
type Patient struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Gender  string `json:"gender"`
    Contact string `json:"contact"`
}
