package dbaccess

import (
	"database/sql"
	"fmt"
	entities "trade-platform/Entities"
)

func FindUserById(id int) entities.UserCredentials {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Query("select * from users where id=$1", id)
	var p entities.UserCredentials
	if result != nil {
		for result.Next() {
			err := result.Scan(&p.Id, &p.Login, &p.Password)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return p
}

func FindUserByLogin(login string) entities.UserCredentials{
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Query("select * from users where login=$1", login)
	var p entities.UserCredentials
	if result != nil {
		for result.Next() {
			err := result.Scan(&p.Id, &p.Login, &p.Password)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return p
}

func FindUserByLoginAndPassword(user entities.AuthRequestJson) entities.UserCredentials {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Query("select * from users where login=$1 and password=$2", user.Login, user.Password)
	var p entities.UserCredentials
	if result != nil {
		for result.Next() {
			err := result.Scan(&p.Id, &p.Login, &p.Password)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return p
}

func InsertUser(user entities.AuthRequestJson){
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("insert into users(login, password) values($1, $2)", user.Login, user.Password)
}