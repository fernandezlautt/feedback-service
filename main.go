package main

import (
	"fernandezlautt/feedback-service/db"
	"fernandezlautt/feedback-service/rabbit/consumer"
	"fernandezlautt/feedback-service/rest"
)

func main() {
	db.ConnectDatabase()
	defer db.DisconnectDatabase()
	consumer.Init()
	rest.Init()
}
