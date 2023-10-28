package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type GoogleClaims struct {
	UserID    string `json:"sub"`
	Email     string `json:"email"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Picture   string `json:"picture"`
	jwt.StandardClaims
}

var PUBLIC_KEY_URL = "https://www.googleapis.com/oauth2/v1/certs"
var ISS_1 = "accounts.google.com"
var ISS_2 = "https://accounts.google.com"

func getGooglePublicKey(keyID string) (string, error) {
	res, err := http.Get(PUBLIC_KEY_URL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	pem := map[string]string{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pem)
	if err != nil {
		return "", err
	}

	key, ok := pem[keyID]
	if !ok {
		return "", errors.New("key not found")
	}

	return key, nil
}

func ValidateGoogleJWT(tokenString string) (GoogleClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &GoogleClaims{}, func(t *jwt.Token) (interface{}, error) {
		pem, err := getGooglePublicKey(t.Header["kid"].(string))
		if err != nil {
			return nil, err
		}

		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, err
		}

		return key, nil
	})
	if err != nil {
		return GoogleClaims{}, err
	}

	claims, ok := token.Claims.(*GoogleClaims)
	if !ok {
		return GoogleClaims{}, errors.New("google jwt is invalid")
	}

	err = godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file %v", err)
		return GoogleClaims{}, errors.New(err.Error())
	}

	CLIENT_ID := os.Getenv("CLIENT_ID")
	if claims.Audience != CLIENT_ID {
		return GoogleClaims{}, errors.New("aud is invalid")
	}

	if claims.Issuer != ISS_1 && claims.Issuer != ISS_2 {
		return GoogleClaims{}, errors.New("iss is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return GoogleClaims{}, errors.New("iss is invalid")
	}

	return *claims, nil
}
