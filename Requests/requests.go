package requests

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	dbaccess "payment-service/DBAccess"
	entities "payment-service/Entities"
	service "payment-service/Service"
	"time"
)

var key, _ = uuid.NewUUID()

func GetPayment(w http.ResponseWriter, r *http.Request) {
	var paymentResponse entities.PaymentSession
	json.NewDecoder(r.Body).Decode(&paymentResponse)
	payment := dbaccess.GetPayment(paymentResponse.SessionId)
	js, err := json.Marshal(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
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
}

func GetPaymentsInPeriod(w http.ResponseWriter, r *http.Request) {
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
}

func ValidateCard(w http.ResponseWriter, r *http.Request) {
	var cardData entities.CardData
	json.NewDecoder(r.Body).Decode(&cardData)
	var response entities.CardValidationResponse
	if service.SimpleLuhnCheck(cardData.Number) {
		payment := dbaccess.GetPayment(cardData.SessionId)
		if payment.ExpireTime > time.Now().String() {
			response.Error = ""
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
}

func HandleRequests() {
	fmt.Println("Server started successfully. Here's admin key:")
	fmt.Println(key.String())
	router := mux.NewRouter()



	getRouter := router.Methods(http.MethodGet).Subrouter()
	postRouter := router.Methods(http.MethodPost).Subrouter()

	getRouter.HandleFunc("/payment", GetPayment)
	getRouter.HandleFunc("/payments", GetPaymentsInPeriod)

	postRouter.HandleFunc("/payment", CreatePayment)
	postRouter.HandleFunc("/validate", ValidateCard)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}
