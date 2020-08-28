package requests

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
	dbaccess "trade-platform/DBAccess"
	entities "trade-platform/Entities"
	service "trade-platform/Service"
)

func GetPayment(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	if token != nil {
		if service.CheckToken(token.Value) {
			var paymentResponse entities.PaymentSession
			json.NewDecoder(r.Body).Decode(&paymentResponse)
			payment := dbaccess.GetPayment(paymentResponse.SessionId)
			js, err := json.Marshal(payment)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(js)
			return
		}
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	if token != nil {
		if service.CheckToken(token.Value) {
			var payment entities.Payment
			json.NewDecoder(r.Body).Decode(&payment)
			id, err := uuid.NewUUID()
			response := entities.PaymentSession{SessionId: id.String()}
			js, err := json.Marshal(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			dbaccess.InsertPayment(payment, id.String(), time.Now().Format("02-01-2006 15:04:05"),
				time.Now().AddDate(0, 0, 7).Format("02-01-2006 15:04:05"))
			w.Write(js)
			return
		}
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func GetPaymentsInPeriod(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	if token != nil {
		if service.CheckToken(token.Value) {
			var period entities.Period
			json.NewDecoder(r.Body).Decode(&period)
			if r.Header.Get("Authorization") != key.String() {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			response := dbaccess.GetPaymentsInPeriod(period.From, period.To)
			js, err := json.Marshal(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(js)
			return
		}
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func ValidateCard(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	if token != nil {
		if service.CheckToken(token.Value) {
			var cardData entities.CardData
			json.NewDecoder(r.Body).Decode(&cardData)
			var response entities.CardValidationResponse
			if service.SimpleLuhnCheck(cardData.Number) {
				payment := dbaccess.GetPayment(cardData.SessionId)
				if payment.ExpireTime > time.Now().String() {
					response.Error = ""
					response.Key = dbaccess.FindProductById(cardData.KeyId).Key
					service.SendEmail()
					service.SendNotificationToOwner()
					dbaccess.MakePaymentComplete(cardData.SessionId, time.Now().Format("02-01-2006 15:04:05"), cardData.Number)
				} else {
					response.Error = "Payment time expired."
				}
			} else {
				response.Error = "Invalid card."
			}
			js, err := json.Marshal(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(js)
			return
		}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
