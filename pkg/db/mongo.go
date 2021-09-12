package db

import (
	"context"
	"go-starter/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB interface {
	ConnectDatabase() *mongo.Database
}

type mongoDB struct {
	db *mongo.Database
}

func NewMongoClient() MongoDB {

	databaseName := config.Get().DbName

	client, err := mongo.NewClient(options.Client().ApplyURI("<MONGODB_URI>"))
	if err != nil {
		log.Fatal("error while creating new monog client. error: ", err)
		return nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("mongo client connection failed. error: ", err)
		return nil
	}

	database := client.Database(databaseName)
	return &mongoDB{
		db: database,
	}
}

func (m mongoDB) ConnectDatabase() *mongo.Database {
	return m.db
}
