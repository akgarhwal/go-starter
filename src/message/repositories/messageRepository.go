package repositories

import (
	"go-starter/infra/db"
	"go-starter/src/message/models"
)

type MessageRepositoryInterface interface {
	Find(string) models.Message
}

type messageRepository struct {
	DB db.MongoDB
}

func NewMessageRepostiory(DB db.MongoDB) MessageRepositoryInterface {
	return &messageRepository{
		DB: DB,
	}
}

func (p *messageRepository) Find(name string) models.Message {
	var message models.Message
	// fetch message from db
	// message = p.DB.ConnectDatabase().Collection(message.CollectionName()).FindOne(nil, bson.M{})

	// dummy message
	message.Name = "Hello " + name
	return message
}
