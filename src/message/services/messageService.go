package services

import (
	"go-starter/src/message/dto"
	"go-starter/src/message/entities"
	"go-starter/src/message/repositories"
)

type MessageService interface {
	GetMessage(dto.MessageRequestBody) entities.Message
}

type messageService struct {
	MessageRepository repositories.MessageRepositoryInterface
}

func NewMessageService(messageRepository repositories.MessageRepositoryInterface) MessageService {
	return &messageService{
		MessageRepository: messageRepository,
	}
}

func (repo *messageService) GetMessage(request dto.MessageRequestBody) entities.Message {
	return repo.MessageRepository.Find(request.Name)
}
