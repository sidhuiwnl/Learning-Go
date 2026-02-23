package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB


func InitDB(){

	var err error

	DB, err = sql.Open("sqlite","./todos.db")

	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS todos(
		id integer primary key autoincrement,
		name TEXT NOT NULL,
		status BOOLEAN NOT NULL DEFAULT 0
	)
	`
	_, err = DB.Exec(createTable)


	if err != nil {
		log.Fatal(err)
	}
}