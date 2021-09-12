package services

import (
	"go-starter/src/message/dto"
	"go-starter/src/message/models"
	"go-starter/src/message/repositories"
)

type MessageService interface {
	GetMessage(dto.MessageRequestBody) models.Message
}

type messageService struct {
	MessageRepository repositories.MessageRepositoryInterface
}

func NewMessageService(messageRepository repositories.MessageRepositoryInterface) MessageService {
	return &messageService{
		MessageRepository: messageRepository,
	}
}

func (repo *messageService) GetMessage(request dto.MessageRequestBody) models.Message {
	return repo.MessageRepository.Find(request.Name)
}
