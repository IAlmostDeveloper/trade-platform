package main

import (
	dbaccess "payment-service/DBAccess"
	requests "payment-service/Requests"
)

func main() {
	dbaccess.CreateDB()
	requests.HandleRequests()
}
