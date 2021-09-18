package main

import (
	"go-starter/library/logger"
	"go-starter/server"
	"go-starter/storage/db"
)

func main() {

	database := db.NewMongoClient()

	server := server.NewServer()
	server.SetDatabase(database)

	logger.Log.Fatal(server.StartHttpServer())
}
