package services

import (
	"database/sql"
	"errors"
	"lottery-api/internal/config"
	"lottery-api/internal/db"
	"lottery-api/internal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define the JWT secret key
var jwtSecret = config.JWT_SECRET

// CreateUser creates a new user in the database
func CreateUser(user *models.User) error {
	_, err := db.Conn.Exec(`
		INSERT INTO users (username, password_hash, role, wallet_balance)
		VALUES ($1, $2, $3, $4)
	`, user.Username, user.PasswordHash, user.Role, user.WalletBalance)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(username string) (*models.User, error) {
	row := db.Conn.QueryRow(`
		SELECT id, username, password_hash, role, wallet_balance, created_at
		FROM users
		WHERE username = $1
	`, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.WalletBalance, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// ValidateJWT parses and validates the JWT token, returning the associated user
func ValidateJWT(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	var user models.User
	err = db.Conn.QueryRow(`
		SELECT id, username, role
		FROM users
		WHERE username = $1
	`, username).Scan(&user.ID, &user.Username, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GenerateJWT generates a JWT token for a given user
func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
