package main

import (
	dbaccess "trade-platform/DBAccess"
	requests "trade-platform/Requests"
)

func main() {
	dbaccess.CreateDB()
	requests.HandleRequests()
}
