package main

import (
	"fernandezlautt/feedback-service/db"
	"fernandezlautt/feedback-service/rest"
)

func main() {
	db.ConnectDatabase()
	defer db.DisconnectDatabase()
	rest.Init()
}
