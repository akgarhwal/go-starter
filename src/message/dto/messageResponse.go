package dto

import "go-starter/src/message/entities"

type MessageResponseBody struct {
	Name string `json:"name"`
}

func ParseFromEntity(data entities.Message) MessageResponseBody {
	return MessageResponseBody{
		Name: data.Name,
	}
}
