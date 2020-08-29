package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"net/http"
	"time"
	configs "trade-platform/Configs"
	entities "trade-platform/Entities"
)

func CreateToken(login string, email string, expirationTime time.Time) (string, error) {
	claims := entities.Claims{
		Login: login,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(configs.JwtSecretKey)
	return tokenString, err
}

func GetUserDataFromToken(tokenString string) (string, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return configs.JwtSecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		login := fmt.Sprintf("%v", claims["login"])
		email := fmt.Sprintf("%v", claims["email"])
		return login, email
	} else {
		fmt.Println(err)
	}
	return "", ""
}

func WriteToken(token string){
	rdb := redis.NewClient(&redis.Options{
		Addr:     configs.RedisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rdb.Set(configs.RedisContext, token, "Ok",
		time.Duration(1000000000*configs.TokenTTLSeconds)) // 10 seconds
	rdb.Save(configs.RedisContext)
}

func CheckToken(token string) bool {
	rdb := redis.NewClient(&redis.Options{
		Addr:     configs.RedisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	result := rdb.Get(configs.RedisContext,token)
	return  result.Val() == "Ok"
}

func AuthorizeUser(token *http.Cookie, err error) bool{
	if token != nil {
		return CheckToken(token.Value)
	}
	return false
}
