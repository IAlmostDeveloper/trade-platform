package requests

import (
	"encoding/json"
	"net/http"
	"time"
	entities "trade-platform/Entities"
	service "trade-platform/Service"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var request entities.AuthRequestJson
	json.NewDecoder(r.Body).Decode(&request)
	isSuccess, user := service.Authenticate(request)
	if !isSuccess{
		http.Error(w, "Incorrect user data.", http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	tokenString, err := service.CreateToken(user.Login, user.Email, expirationTime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	service.WriteToken(tokenString)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user entities.RegRequestJson
	json.NewDecoder(r.Body).Decode(&user)
	if !service.Register(user){
		http.Error(w, "User with this name already exists", http.StatusBadRequest)
	}
}
