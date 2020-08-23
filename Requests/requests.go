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
	"time"
	dbaccess "trade-platform/DBAccess"
	entities "trade-platform/Entities"
	service "trade-platform/Service"
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

func GetProducts(w http.ResponseWriter, r *http.Request){
	response := dbaccess.GetAllProducts()
	//fmt.Println(r.Cookie("token"))
	//service.WriteToken("1234")
	token, err := r.Cookie("token")
	if token != nil{
		if !service.CheckToken(token.String()){
			http.Error(w, "Authorization failed", http.StatusUnauthorized)
			return
		}
	}
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := service.GetIdFromPath(r.URL.Path)
	response := dbaccess.FindProductById(id)
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	if token != nil{
		var result = service.GetLoginFromToken(token.Value)
		w.Write([]byte(result))
	}
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var request entities.AuthRequestJson
	json.NewDecoder(r.Body).Decode(&request)
	user := dbaccess.FindUserByLoginAndPassword(request)
	if user.Id == 0 {
		http.Error(w, "Incorrect user data.", http.StatusUnauthorized)
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	tokenString, err := service.CreateToken(user.Login, expirationTime)
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
	var user entities.AuthRequestJson
	json.NewDecoder(r.Body).Decode(&user)
	if dbaccess.FindUserByLogin(user.Login).Id==0{
		dbaccess.InsertUser(user)
		return
	}
	http.Error(w, "User with this name already exists", http.StatusBadRequest)

}

func HandleRequests() {
	fmt.Println("Server started successfully. Here's admin key:")
	fmt.Println(key.String())
	router := mux.NewRouter()



	getRouter := router.Methods(http.MethodGet).Subrouter()
	postRouter := router.Methods(http.MethodPost).Subrouter()

	getRouter.HandleFunc("/payment", GetPayment)
	getRouter.HandleFunc("/payments", GetPaymentsInPeriod)

	getRouter.HandleFunc("/products", GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", GetProduct)


	postRouter.HandleFunc("/payment", CreatePayment)
	postRouter.HandleFunc("/validate", ValidateCard)

	postRouter.HandleFunc("/products", CreateProduct)

	postRouter.HandleFunc("/auth", Authenticate)
	postRouter.HandleFunc("/register", Register)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}