package main

import (
	_ "modernc.org/sqlite"
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB(){

	var err error

	DB, err = sql.Open("sqlite","./todos.db")

	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	 CREATE TABLE IF NOT EXISTS url (
	 	id integer primary key autoincrement,
		short_code text unique not null,
		original_url text not null

	 )
	`

	_, err = DB.Exec(createTable)

	if err != nil {
		log.Fatal(err)
	}
}

