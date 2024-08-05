package services

import (
	"errors"
	"lottery-api/internal/db"
	"lottery-api/internal/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(req models.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("INSERT INTO general_users (username, password_hash) VALUES ($1, $2)", req.Username, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func LoginUser(req models.LoginRequest) (string, error) {
	var hashedPassword string
	err := db.DB.QueryRow("SELECT password_hash FROM general_users WHERE username = $1", req.Username).Scan(&hashedPassword)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := generateToken(req.Username)
	return token, nil
}

func generateToken(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	tokenString, _ := token.SignedString([]byte("secret"))
	return tokenString
}
