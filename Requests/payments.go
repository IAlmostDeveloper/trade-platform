package requests

import (
	"encoding/json"
	"net/http"
	entities "trade-platform/Entities"
	service "trade-platform/Service"
)

func GetPayment(w http.ResponseWriter, r *http.Request) {
	if service.AuthorizeUser(r.Cookie("token")) {
		var paymentResponse entities.PaymentSession
		json.NewDecoder(r.Body).Decode(&paymentResponse)
		js := service.GetPayment(paymentResponse)
		if js == nil{
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write(js)
		return
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	if service.AuthorizeUser(r.Cookie("token")) {
		var payment entities.Payment
		json.NewDecoder(r.Body).Decode(&payment)
		js := service.CreatePayment(payment)
		if js == nil{
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		w.Write(js)
		return
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func GetPaymentsInPeriod(w http.ResponseWriter, r *http.Request) {
	if service.AuthorizeUser(r.Cookie("token")) {
		var period entities.Period
		json.NewDecoder(r.Body).Decode(&period)
		js := service.GetPaymentsInPeriod(period)
		if js == nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		w.Write(js)
		return
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func CompletePayment(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if service.AuthorizeUser(token, err) {
		var cardData entities.CardData
		json.NewDecoder(r.Body).Decode(&cardData)
		customerLogin, customerEmail := service.GetLoginAndEmailFromToken(token.Value)
		js := service.CompletePayment(cardData, customerLogin, customerEmail)
		if js == nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(js)
		return
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
