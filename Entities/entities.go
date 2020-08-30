package entities

import "github.com/dgrijalva/jwt-go"

type Payment struct {
	Sum     int    `json:"sum"`
	Purpose string `json:"purpose"`
	KeyId int `json:"key_id"`
}

type PaymentFromDB struct {
	Id            int    `json:"id"`
	Sum           int    `json:"sum"`
	Purpose       string `json:"purpose"`
	KeyId int `json:"key_id"`
	SessionId     string `json:"session_id"`
	CreatedTime   int64 `json:"created_time"`
	CompletedTime int64 `json:"completed_time"`
	ExpireTime    int64 `json:"expire_time"`
	Completed     bool   `json:"completed"`
	Card          string `json:"card"`
}

type PaymentSession struct {
	SessionId string `json:"session_id"`
}

type Period struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type CardData struct {
	User       string `json:"user"`
	Number     string `json:"number"`
	CVV        int    `json:"cvv"`
	ExpireDate string `json:"expire_date"`
	SessionId  string `json:"session_id"`
}

type CardValidationResponse struct {
	Error string `json:"error"`
	Key string `json:"key"`
}

type AuthRequestJson struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type RegRequestJson struct{
	Login string `json:"login"`
	Email string `json:"email"`
	Domain string `json:"domain"`
	Password string `json:"password"`
}

type Claims struct {
	Login string `json:"login"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type UserCredentials struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
	Domain string `json:"domain"`
	Password string `json:"password"`
}

type Product struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Key        string `json:"key"`
	Price      int    `json:"price"`
	Commission int    `json:"commission"`
	Owner      string `json:"owner"`
}

type ProductInfo struct {
	Name string `json:"name"`
}

type KeyIdJson struct {
	KeyId int `json:"key_id"`
}

type Purchase struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Key string `json:"key"`
	Date string `json:"date"`
	Buyer string `json:"buyer"`
	Owner string `json:"owner"`
}

type PurchaseInfo struct {
	GameName string `json:"game_name"`
	Buyer string `json:"buyer"`
	PaymentSessionId string `json:"payment_session_id"`
	CommissionSum float32 `json:"commission_sum"`
	OwnerIncome float32 `json:"owner_income"`
}

type EmailContent struct {
	CustomerEmail string `json:"customer_email"`
	GameKey string `json:"game_key"`
}
