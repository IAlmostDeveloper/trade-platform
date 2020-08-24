package requests

import (
	"encoding/json"
	"net/http"
	"time"
	dbaccess "trade-platform/DBAccess"
	entities "trade-platform/Entities"
	service "trade-platform/Service"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var request entities.AuthRequestJson
	json.NewDecoder(r.Body).Decode(&request)
	user := dbaccess.FindUserByLoginAndPassword(request)
	if user.Id == 0 {
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
	if dbaccess.FindUserByLogin(user.Login).Id==0{
		dbaccess.InsertUser(user)
		return
	}
	http.Error(w, "User with this name already exists", http.StatusBadRequest)
}
