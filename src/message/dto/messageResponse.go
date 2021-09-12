package dto

import "go-starter/src/message/models"

type MessageResponseBody struct {
	Name string `json:"name"`
}

func ParseFromEntity(data models.Message) MessageResponseBody {
	return MessageResponseBody{
		Name: data.Name,
	}
}
