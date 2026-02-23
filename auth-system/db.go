package main


import (
	_ "modernc.org/sqlite"
	"database/sql"
	"log"
)

var DB *sql.DB


func InitDB(){
	
	var err error
	DB, err = sql.Open("sqlite","./auth.db")

	if err != nil {
		log.Fatal(err)
	}

	query := `
	 CREATE TABLE IF NOT EXISTS users (
	  id integer  primary key autoincrement,
	  email text unique not null,
	  password text not null

	 )
	`

	_, err = DB.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

}