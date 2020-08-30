package dbaccess

import (
	"fmt"
	entities "trade-platform/Entities"
)

func GetAllProducts() []entities.Product{
	db := OpenDB()
	defer db.Close()
	result, _ := db.Query("select * from products")
	var products []entities.Product
	if result != nil {
		for result.Next() {
			var p entities.Product
			err := result.Scan(&p.Id, &p.Name, &p.Key, &p.Price, &p.Commission, &p.Owner)
			if err != nil {
				fmt.Println(err)
				continue
			}
			products = append(products, p)
		}
	}
	return products
}

func GetOwnerProducts(owner string) []entities.Product {
	db := OpenDB()
	defer db.Close()
	result, _ := db.Query("select * from products where owner=$1", owner)
	var products []entities.Product
	if result != nil {
		for result.Next() {
			var p entities.Product
			err := result.Scan(&p.Id, &p.Name, &p.Key, &p.Price, &p.Commission, p.Owner)
			if err != nil {
				fmt.Println(err)
				continue
			}
			products = append(products, p)
		}
	}
	return products
}

func FindProductById(id int) entities.Product {
	db := OpenDB()
	defer db.Close()
	result, _ := db.Query("select * from products where id=$1", id)
	var p entities.Product
	if result != nil {
		for result.Next() {
			err := result.Scan(&p.Id, &p.Name, &p.Key, &p.Price, &p.Commission, &p.Owner)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return p
}

func FindProductByName(name string) entities.Product {
	db := OpenDB()
	defer db.Close()
	result, _ := db.Query("select * from products where name=$1 limit 1", name)
	var p entities.Product
	if result != nil {
		for result.Next() {
			err := result.Scan(&p.Id, &p.Name, &p.Key, &p.Price, &p.Commission, &p.Owner)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return p
}

func InsertProduct(product entities.Product){
	db := OpenDB()
	defer db.Close()
	db.Exec("insert into products(name, key, price, commission, owner) " +
		"values($1,$2,$3,$4,$5)", product.Name, product.Key, product.Price, product.Commission, product.Owner)
}

func DeleteProduct(id int){
	db := OpenDB()
	defer db.Close()
	db.Exec("delete from products where id=$1", id)
}
