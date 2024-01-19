package auth

import (
	"Todo-Go/app/model"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your-secret-key")

type Claims struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}

func GenerateJWT(user model.User) (string, error) {
	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*model.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return &claims.User, nil
	}

	return nil, errors.New("Invalid token")
}

func AuthenticateUser(username, password string) (*model.User, error) {
	if username != "" && password != "" {
		return &model.User{
			ID:       1,
			Username: username,
			Email:    "user@example.com",
		}, nil
	}
	return nil, fmt.Errorf("Authentication failed")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
