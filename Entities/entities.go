package entities

import "github.com/dgrijalva/jwt-go"

type Payment struct {
	Sum     int    `json:"sum"`
	Purpose string `json:"purpose"`
}

type PaymentFromDB struct {
	Id            int    `json:"id"`
	Sum           int    `json:"sum"`
	Purpose       string `json:"purpose"`
	SessionId     string `json:"session_id"`
	CreatedTime   string `json:"created_time"`
	CompletedTime string `json:"completed_time"`
	ExpireTime    string `json:"expire_time"`
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
}

type AuthRequestJson struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

type UserCredentials struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Password string `json:"password"`
}

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Key string `json:"key"`
	Price int `json:"price"`
	Commission int `json:"commission"`
	OwnerId int `json:"owner_id"`
}
