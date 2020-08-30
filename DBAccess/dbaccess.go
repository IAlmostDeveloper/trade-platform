package dbaccess

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("create table if not exists users(id integer primary key autoincrement, " +
		"login text unique, email text unique, domain text, password text)")
	db.Exec("create table if not exists payments(id integer primary key autoincrement," +
		"sum integer, purpose text, key_id integer, session_id text, created_time text," +
		"completed_time text, expire_time text, completed numeric, card text)")
	db.Exec("create table if not exists purchases(id integer primary key autoincrement, " +
		"name text, key text unique, date text, buyer text, owner text)")
	db.Exec("create table if not exists products(id integer primary key autoincrement, " +
		"name text, key text unique, price integer, commission integer, owner text)")

	db.Exec("create table if not exists platform(id integer primary key autoincrement, " +
		"key text unique, value text)")
	db.Exec("insert or ignore into platform(key, value) values('balance', '0');")
}

func OpenDB() *sql.DB {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	return db
}
