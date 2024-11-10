package main

import (
    "log"
    "myapp/config"
    "myapp/db"
    "myapp/routes"
    "net/http"
	
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Initialize the database connection
    if err := config.InitDB(); err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Run migrations
    if err := db.Migrate(); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    // Register API routes
    router := routes.RegisterRoutes()

    // Serve static files from the "static" directory
    fs := http.FileServer(http.Dir("./static"))
    router.PathPrefix("/").Handler(fs)

    // Start the server
    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
