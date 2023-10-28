package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mdayat/capstone-rest-service/internal/auth"
)

func Auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	fmt.Println(claims.Picture)
}
