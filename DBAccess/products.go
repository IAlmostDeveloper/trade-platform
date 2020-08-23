package dbaccess

import (
	"database/sql"
	"fmt"
	entities "trade-platform/Entities"
)

func FindProductById(id int) entities.Product {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Query("select * from products where id=$1", id)
	var p entities.Product
	if result != nil {
		for result.Next() {
			err := result.Scan(&p.Id, &p.Name, &p.Key, &p.Price, &p.Commission)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return p
}

func FindProductByName(name string) entities.Product {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Query("select * from products where name=$1 limit 1", name)
	var p entities.Product
	if result != nil {
		for result.Next() {
			err := result.Scan(&p.Id, &p.Name, &p.Key, &p.Price, &p.Commission)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return p
}

func InsertProduct(product entities.Product){
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Query("insert into products(name, key, price, commission) " +
		"values($1,$2,$3,$4)", product.Name, product.Key, product.Price, product.Commission)
}

func DeleteProduct(id int){
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Query("delete from products where id=$1", id)
}
