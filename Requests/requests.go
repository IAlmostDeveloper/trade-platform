package requests

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func HandleRequests() {
	fmt.Println("Server started successfully.")
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
	postRouter.HandleFunc("/purchase", PurchaseProduct)

	postRouter.HandleFunc("/auth", Authenticate)
	postRouter.HandleFunc("/register", Register)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	port := ":" + os.Getenv("PORT")
	fmt.Println(port)

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}