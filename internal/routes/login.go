package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mdayat/capstone-rest-service/internal/auth"
)

type LoginResponse struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	var tokenString string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tokenString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	claims, err := auth.ValidateGoogleJWT(tokenString)
	if err != nil {
		errMsg := "invalid google auth"
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(errMsg))
		return
	}

	refreshToken, err := auth.CreateToken(claims.UserID, "REFRESH")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	accessToken, err := auth.CreateToken(claims.UserID, "ACCESS")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	loginResponse := LoginResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}

	loginResponseJSON, err := json.Marshal(loginResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(loginResponseJSON)
}
