package dbaccess

import (
	"database/sql"
	"fmt"
	entities "trade-platform/Entities"
)

func GetPayment(session_id string) entities.PaymentFromDB {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Query("select * from payments where session_id=$1", session_id)
	var p entities.PaymentFromDB
	if result != nil {
		for result.Next() {
			err := result.Scan(&p.Id, &p.Sum, &p.Purpose, &p.SessionId,
				&p.CreatedTime, &p.CompletedTime, &p.ExpireTime, &p.Completed, &p.Card)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return p
}

func GetPaymentsInPeriod(from string, to string) []entities.PaymentFromDB {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Query("select * from payments where created_time>=$1 and created_time<=$2", from, to)
	var payments []entities.PaymentFromDB
	for result.Next() {
		var p entities.PaymentFromDB
		err := result.Scan(&p.Id, &p.Sum, &p.Purpose, &p.SessionId,
			&p.CreatedTime, &p.CompletedTime, &p.ExpireTime, &p.Completed, &p.Card)
		if err != nil {
			fmt.Println(err)
			continue
		}
		payments = append(payments, p)
	}
	return payments
}

func InsertPayment(payment entities.Payment, session_id string, created_time string, expire_time string) {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("insert into payments(sum, purpose, session_id, created_time," +
		"completed_time, expire_time, completed, card) "+
		"values($1, $2, $3, $4, '', $5, false, '')",
		payment.Sum, payment.Purpose, session_id, created_time, expire_time)
}

func MakePaymentComplete(session_id string, completed_time string, card_number string) {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("update payments set completed=true, completed_time=$1, card=$2 where session_id=$3",
		completed_time, card_number, session_id)
}
