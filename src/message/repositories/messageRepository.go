package repositories

import (
	"go-starter/pkg/db"
	"go-starter/src/message/entities"
)

type MessageRepositoryInterface interface {
	Find(string) entities.Message
}

type messageRepository struct {
	DB db.MongoDB
}

func NewMessageRepostiory(DB db.MongoDB) MessageRepositoryInterface {
	return &messageRepository{
		DB: DB,
	}
}

func (p *messageRepository) Find(name string) entities.Message {
	var message entities.Message
	// fetch message from db
	// message = p.DB.ConnectDatabase().Collection(message.CollectionName()).FindOne(nil, bson.M{})

	// dummy message
	message.Name = "Hello " + name
	return message
}
