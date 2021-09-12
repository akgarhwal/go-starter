package main

import (
	"go-starter/infra/db"
	"go-starter/internal/logger"
	"go-starter/internal/server"
)

func main() {

	log := logger.Log
	database := db.NewMongoClient()

	server := server.NewServer()
	server.SetDatabase(database)
	server.SetLogger(log)

	log.Fatal(server.StartHttpServer())
}
