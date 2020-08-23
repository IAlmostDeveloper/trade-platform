package tests

import (
	"github.com/magiconair/properties/assert"
	service "payment-service/Service"
	"testing"
)

func TestLuhnCheck_RightCard(t *testing.T) {
	cardNumber := "4561261212345467"
	assert.Equal(t, service.SimpleLuhnCheck(cardNumber), true)
}

func TestLuhnCheck_WrongCard(t *testing.T) {
	cardNumber := "4561261212345464"
	assert.Equal(t, service.SimpleLuhnCheck(cardNumber), false)
}

func TestLuhnCheck_WrongNumbersCountFails(t *testing.T) {
	cardNumber := "45612611"
	assert.Equal(t, service.SimpleLuhnCheck(cardNumber), false)
}

func TestLuhnCheck_NotNumberStringFails(t *testing.T) {
	cardNumber := "Not a number!!!!"
	assert.Equal(t, service.SimpleLuhnCheck(cardNumber), false)
}
