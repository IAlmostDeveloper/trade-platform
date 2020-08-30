package dbaccess

func AddPaymentCommission(commissionSum float32){
	db := OpenDB()
	defer db.Close()
	db.Exec("update platform set value=value + $1 where key='balance'", commissionSum)
}
