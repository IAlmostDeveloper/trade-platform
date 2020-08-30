package service

import (
	"encoding/json"
	"github.com/google/uuid"
	service "payment-service/Service"
	"time"
	configs "trade-platform/Configs"
	dbaccess "trade-platform/DBAccess"
	entities "trade-platform/Entities"
)

func GetPayment(session entities.PaymentSession) []byte {
	payment := dbaccess.GetPayment(session.SessionId)
	js, err := json.Marshal(payment)
	if err != nil {
		return nil
	}
	return js
}

func CreatePayment(payment entities.Payment) []byte {
	id, _ := uuid.NewUUID()
	response := entities.PaymentSession{SessionId: id.String()}
	dbaccess.InsertPayment(payment, id.String(), time.Now(),
		time.Now().AddDate(0, 0, 7))
	js, err := json.Marshal(response)
	if err != nil {
		return nil
	}
	return js
}

func GetPaymentsInPeriod(period entities.Period) []byte {
	from, _ := time.Parse(configs.DateTimeLayout, period.From)
	to, _ := time.Parse(configs.DateTimeLayout, period.To)
	response := dbaccess.GetPaymentsInPeriod(from, to)
	js, err := json.Marshal(response)
	if err != nil {
		return nil
	}
	return js
}

func ValidateCard(cardData entities.CardData, customerLogin string, customerEmail string) []byte {
	var response entities.CardValidationResponse
	if service.SimpleLuhnCheck(cardData.Number) {
		payment := dbaccess.GetPayment(cardData.SessionId)

		if PaymentNotExpired(payment.ExpireTime) {
			product := dbaccess.FindProductById(payment.KeyId)
			dbaccess.MakePaymentComplete(cardData.SessionId, time.Now(), cardData.Number)
			dbaccess.DeleteProduct(product.Id)

			commissionSum := SendCommissionToPlatform(product)

			purchaseInfo := entities.PurchaseInfo{GameName: product.Name, Buyer: customerLogin,
				PaymentSessionId: payment.SessionId, CommissionSum: commissionSum,
				OwnerIncome: float32(product.Price) - commissionSum,
			}

			SendEmail(customerEmail, response.Key)

			domain := dbaccess.FindUserByLogin(product.Owner).Domain
			SendNotificationToOwner(domain, purchaseInfo)

			response = entities.CardValidationResponse{Error: "", Key: product.Key}
		} else {
			response.Error = "Payment time expired."
		}
	} else {
		response.Error = "Invalid card."
	}
	js, err := json.Marshal(response)
	if err != nil {
		return nil
	}
	return js
}
