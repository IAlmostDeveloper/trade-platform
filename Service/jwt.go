package service

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"time"
	entities "trade-platform/Entities"
)

var JwtKey = []byte("3059a5e0-e543-11ea-9af4-b4b52f893c01")
var ctx = context.Background()

func CreateToken(login string, email string, expirationTime time.Time) (string, error) {
	claims := entities.Claims{
		Login: login,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	return tokenString, err
}

func GetUserDataFromToken(tokenString string) (string, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return JwtKey, nil
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
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rdb.Set(ctx, token, "Ok", 1000000000 * 60) // 10 seconds
	rdb.Save(ctx)
}

func CheckToken(token string) bool {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	result := rdb.Get(ctx,token)
	return  result.Val() == "Ok"
}
