package tests

import (
	"encoding/json"
	"fmt"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	dbaccess "payment-service/DBAccess"
	entities "payment-service/Entities"
	requests "payment-service/Requests"
	"strings"
	"testing"
)

type Session struct {
	SessionId string `json:"session_id"`
}

func TestCreateAndGetPaymentRequests(t *testing.T) {
	dbaccess.CreateDB()
	req, err := http.NewRequest("POST", "/payment", strings.NewReader(`{"sum":50000, "purpose":"Example"}`))
	req.Header.Set("content-type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	postHandler := http.HandlerFunc(requests.CreatePayment)
	postHandler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("postHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var response Session
	json.NewDecoder(rr.Body).Decode(&response)
	req, err = http.NewRequest("GET", "/payment", strings.NewReader(fmt.Sprintf(`{"session_id": "%s"}`, response.SessionId)))
	req.Header.Set("content-type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	getHandler := http.HandlerFunc(requests.GetPayment)
	getHandler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("postHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var payment entities.PaymentFromDB
	json.NewDecoder(rr.Body).Decode(&payment)

	assert.Equal(t, payment.Sum, 50000)
	assert.Equal(t, payment.Purpose, "Example")
}
