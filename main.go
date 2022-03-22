package main

import (
	"github.com/renatodaltiba/rest-api-go/database"
	"github.com/renatodaltiba/rest-api-go/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()

}
