package tests

import (
	"github.com/magiconair/properties/assert"
	dbaccess "payment-service/DBAccess"
	entities "payment-service/Entities"
	"testing"
)

func TestInsertAndGetPaymentFromDB(t *testing.T) {
	dbaccess.CreateDB()
	dbaccess.InsertPayment(entities.Payment{Sum: 50000, Purpose: "Example"},
		"12345", "10-10-2020", "20-20-2020")
	payment := dbaccess.GetPayment("12345")
	assert.Equal(t, payment.Sum, 50000)
	assert.Equal(t, payment.Purpose, "Example")
}
