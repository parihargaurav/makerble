package config

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq" // PostgreSQL driver
)

// DB is the global variable to hold the database connection
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
    // Load the database credentials from environment variables
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Define the PostgreSQL connection string
    connStr := "user=" + user + " password=" + password + " dbname=" + dbName + " sslmode=disable search_path=public"

    // Open a connection to the database
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        return err
    }

    // Test the database connection
    if err := DB.Ping(); err != nil {
        return err
    }

    log.Println("Successfully connected to the database.")
    return nil
}
