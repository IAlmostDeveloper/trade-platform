package requests

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
	configs "trade-platform/Configs"
	dbaccess "trade-platform/DBAccess"
	entities "trade-platform/Entities"
	service "trade-platform/Service"
)

func GetPayment(w http.ResponseWriter, r *http.Request) {
	if service.AuthorizeUser(r.Cookie("token")) {
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
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	if service.AuthorizeUser(r.Cookie("token")) {
		var payment entities.Payment
		json.NewDecoder(r.Body).Decode(&payment)
		id, err := uuid.NewUUID()
		response := entities.PaymentSession{SessionId: id.String()}
		js, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dbaccess.InsertPayment(payment, id.String(), time.Now().Format(configs.DateTimeLayout),
			time.Now().AddDate(0, 0, 7).Format(configs.DateTimeLayout))
		w.Write(js)
		return
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func GetPaymentsInPeriod(w http.ResponseWriter, r *http.Request) {
	if service.AuthorizeUser(r.Cookie("token")) {
		var period entities.Period
		json.NewDecoder(r.Body).Decode(&period)
		response := dbaccess.GetPaymentsInPeriod(period.From, period.To)
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

func ValidateCard(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if service.AuthorizeUser(token, err) {
		var cardData entities.CardData
		var response entities.CardValidationResponse
		json.NewDecoder(r.Body).Decode(&cardData)
		if service.SimpleLuhnCheck(cardData.Number) {
			payment := dbaccess.GetPayment(cardData.SessionId)
			if service.PaymentNotExpired(payment.ExpireTime) {
				product := dbaccess.FindProductById(payment.KeyId)
				customerLogin, customerEmail := service.GetLoginAndEmailFromToken(token.Value)
				dbaccess.MakePaymentComplete(cardData.SessionId, time.Now().Format(configs.DateTimeLayout), cardData.Number)
				dbaccess.DeleteProduct(product.Id)
				
				commissionSum := service.SendCommissionToPlatform(product)
				purchaseInfo := entities.PurchaseInfo{GameName: product.Name, Buyer: customerLogin,
					PaymentSessionId: payment.SessionId, CommissionSum: commissionSum,
				}
				service.SendEmail(customerEmail, response.Key)
				service.SendNotificationToOwner(purchaseInfo)

				response = entities.CardValidationResponse{Error: "", Key: product.Key}
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
