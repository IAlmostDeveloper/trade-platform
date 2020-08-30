package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
	dbaccess "trade-platform/DBAccess"
	entities "trade-platform/Entities"
)

func SimpleLuhnCheck(cardNumber string) bool {
	if len(cardNumber) != 16 {
		return false
	}
	a := strings.Split(cardNumber, "")
	sum := 0
	for i, s := range a {
		num, _ := strconv.Atoi(s)
		if i%2 == 0 {
			if 2*num > 9 {
				sum += 2*num - 9
			} else {
				sum += 2 * num
			}
		} else {
			sum += num
		}
	}
	return sum%10 == 0 && sum > 0
}

func GetIdFromPath(path string) int {
	p := strings.Split(path, "/")
	res, _ := strconv.Atoi(p[2])
	return res
}

func PaymentNotExpired(expireTime int64) bool{
	return expireTime > time.Now().Unix()
}

func SendNotificationToOwner(domain string, purchaseInfo entities.PurchaseInfo) {
	requestBody, _ := json.Marshal(purchaseInfo)
	http.Post(domain, "application/json", bytes.NewBuffer(requestBody))
}

func SendCommissionToPlatform(commissionSum float32){
	dbaccess.AddPaymentCommission(commissionSum)
}
