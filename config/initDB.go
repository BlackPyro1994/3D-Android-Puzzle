package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

// Db handle
var Db *sql.DB
var err error

func InitSQLiteDB() {

	fmt.Println("Initialize SQLite Database")

	Db, err = sql.Open("sqlite3", "./data/borgdirmedia")

	if err != nil {
		fmt.Println("FEHLER")
		panic(err)
	} else {
		fmt.Println("Erfolgreich Verbindung mit Datenbank aufgebaut !")
	}
}

func InitPostgresDB() {

	fmt.Println("Initialize Postgres Database")

	connStr := "user=borgdirmedia dbname=borgdirmedia password=borgdirmedia host=localhost port=5431 sslmode=disable"
	Db, err = sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("FEHLER")
		log.Fatal(err)
	} else {
		fmt.Println("Erfolgreich Verbindung mit Datenbank aufgebaut !")
	}
}
