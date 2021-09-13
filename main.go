package main

import (
	"go-starter/infra/db"
	"go-starter/internal/logger"
	"go-starter/internal/server"
)

func main() {

	database := db.NewMongoClient()

	server := server.NewServer()
	server.SetDatabase(database)

	logger.Log.Fatal(server.StartHttpServer())
}
