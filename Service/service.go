package service

import (
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
	"time"
	configs "trade-platform/Configs"
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

func PaymentNotExpired(expireTime string) bool{
	now := time.Now()
	expire, _ := time.Parse(configs.DateTimeLayout, expireTime)
	return expire.After(now)
}

func SendEmail(customerEmail string, key string) {
	auth := smtp.PlainAuth("", configs.SmtpClientEmail, "password", configs.SmtpClientHost)
	to := []string{customerEmail}
	msg := []byte("To: "+ customerEmail + "\r\n" +
		"Subject: Trade platform!\r\n" +
		"\r\n" +
		"Thanks for your purchase! Here's your key: " + key + "\r\n")
	err := smtp.SendMail(configs.SmtpClientAddress, auth, configs.SmtpClientEmail, to, msg)
	if err != nil {
		fmt.Println(err)
	}
}

func SendNotificationToOwner(purchaseInfo entities.PurchaseInfo) {

}

func SendCommissionToPlatform(product entities.Product) float32{
	commissionSum := float32(product.Price) * float32(product.Commission) / 100
	dbaccess.AddPaymentCommission(commissionSum)
	return commissionSum
}
