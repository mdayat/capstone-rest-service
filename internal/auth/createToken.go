package auth

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func signToken(signMethod *jwt.SigningMethodHMAC, claims *jwt.RegisteredClaims) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: (%v)", err)
		return "", err
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(signMethod, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Printf("Failed to encode JWT: (%v)", err)
		return "", err
	}

	return signedToken, nil
}

func CreateToken(userID string, tokenTypes string) (string, error) {
	var expireTime *jwt.NumericDate
	var issuer string

	if tokenTypes == "REFRESH" {
		issuer = "edumanagerpro-refresh"
		daysInHours := 360
		expireTime = jwt.NewNumericDate(time.Now().UTC().Add(time.Minute * time.Duration(daysInHours)))
	} else if tokenTypes == "ACCESS" {
		issuer = "edumanagerpro-access"
		minutes := 10
		expireTime = jwt.NewNumericDate(time.Now().UTC().Add(time.Minute * time.Duration(minutes)))
	} else {
		errMsg := "incorrect token types. Valid token types are only \"REFRESH\" and \"ACCESS\""
		log.Printf("Incorrect token types %v", tokenTypes)
		return "", errors.New(errMsg)
	}

	claims := &jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   userID,
		Audience:  jwt.ClaimStrings{"http://localhost:3000"},
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: expireTime,
	}

	token, err := signToken(jwt.SigningMethodHS256, claims)
	if err != nil {
		log.Printf("Failed to sign JWT: (%v)", err)
		return "", err
	}

	return token, nil
}
