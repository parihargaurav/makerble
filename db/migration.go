package db

import (
    "log"
    "myapp/config"
)

func Migrate() error {
    // Create "patients" table if it doesn't exist
    _, err := config.DB.Exec(`
        CREATE TABLE IF NOT EXISTS patients (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100),
            age INT,
            gender VARCHAR(10),
            contact VARCHAR(20)
        );
    `)
    if err != nil {
        return err
    }

    // Create "users" table if it doesn't exist
    _, err = config.DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(255) NOT NULL UNIQUE,
            password VARCHAR(255) NOT NULL,
            role VARCHAR(50) NOT NULL CHECK (role IN ('doctor', 'receptionist'))
        );
    `)
    if err != nil {
        return err
    }

    log.Println("Migration completed successfully.")
    return nil
}
