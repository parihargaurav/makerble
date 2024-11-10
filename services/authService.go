package services

import (
    "database/sql"
    "errors"
    "fmt"
    "myapp/config"
    "myapp/models"
    "github.com/dgrijalva/jwt-go"
    "time"
)

// Authenticate checks the database for username and password match and returns a JWT token and user
func Authenticate(username, password string) (*models.User, string, error) {
    var user models.User

    // Query the database for user credentials
    err := config.DB.QueryRow("SELECT id, username, password, role FROM users WHERE username = $1 AND password = $2", username, password).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
    if err == sql.ErrNoRows {
        return nil, "", errors.New("invalid credentials")
    } else if err != nil {
        return nil, "", fmt.Errorf("database error: %v", err)
    }

    // Generate JWT token
    token, err := GenerateJWT(&user)
    if err != nil {
        return nil, "", fmt.Errorf("failed to generate token: %v", err)
    }

    return &user, token, nil
}

// GenerateJWT generates a JWT token for the user
func GenerateJWT(user *models.User) (string, error) {
    claims := jwt.MapClaims{
        "username": user.Username,
        "role":     user.Role,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte("JWT_SECRET")) // Replace with your secret key
    if err != nil {
        return "", errors.New("failed to generate token")
    }

    return signedToken, nil
}
