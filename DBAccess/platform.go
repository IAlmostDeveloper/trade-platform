package dbaccess

import "database/sql"

func AddPaymentCommission(commissionSum float32){
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("update platform set value=value + $1 where key='balance'", commissionSum)
}
