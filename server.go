package main

import (
	"net/http"
	"./app/route"
	"fmt"
	"./config"
)

func main() {

	fmt.Println("Start")

	// config.InitSQLiteDB()
	config.InitPostgresDB()

	route.Handler()

	fs := http.FileServer(http.Dir("./"))

	http.Handle("/static/", fs)

	http.ListenAndServe(":80", nil)

	fmt.Println("Exit")
}
