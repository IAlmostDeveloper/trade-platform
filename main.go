package main

import (
	"sync"
	dbaccess "trade-platform/DBAccess"
	"trade-platform/EmailSender"
	requests "trade-platform/Requests"
)

func main() {
	go func(){
		dbaccess.CreateDB()
		requests.HandleRequests()
	}()
	go func(){
		EmailSender.Start()
	}()

	//Freeze main goroutine by locking locked mutex
	mutex := new(sync.Mutex)
	mutex.Lock()
	mutex.Lock()
}
