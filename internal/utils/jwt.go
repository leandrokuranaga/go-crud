package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func loadEnv() string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return os.Getenv("SECRET_KEY")
}

func GenerateToken(email string, userId int64) (string, error) {
	secretKey := loadEnv()

	if secretKey == "" {
		log.Fatal("The SECRET_KEY environment variable is not defined or empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {

	secretKey := loadEnv()

	if secretKey == "" {
		log.Fatal("The SECRET_KEY environment variable is not defined or empty")
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return 0, errors.New("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Invalid token!")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid Token Claims")
	}

	var userId int64

	if userIdFloat, ok := claims["userId"].(float64); ok {
    userId = int64(userIdFloat)
    return userId, nil
	} else {
		return 0, errors.New("userId claim not found")
	}
}
