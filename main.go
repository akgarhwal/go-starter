package main

import (
	"go-starter/internal/server"
	"go-starter/pkg/db"
	"go-starter/pkg/logger"
)

func main() {

	log := logger.Log
	database := db.NewMongoClient()

	server := server.NewServer()
	server.SetDatabase(database)
	server.SetLogger(log)

	log.Fatal(server.StartHttpServer())
}
