package handlers

import (
	"encoding/json"
	"go-starter/internal/response"
	"go-starter/src/message/dto"
	"go-starter/src/message/services"
	"net/http"
)

type MessageHandlers interface {
	GetMessage(http.ResponseWriter, *http.Request)
}

type messageHandlers struct {
	MessageService services.MessageService
}

func NewHttpHandler(messageService services.MessageService) MessageHandlers {
	return &messageHandlers{
		MessageService: messageService,
	}
}

func (services *messageHandlers) GetMessage(rw http.ResponseWriter, r *http.Request) {

	var messageRequest dto.MessageRequestBody
	err := json.NewDecoder(r.Body).Decode(&messageRequest)
	if err != nil {
		//http.Error(rw, "invalid request body", http.StatusBadRequest)
		response.JsonResponse(rw, http.StatusBadRequest, "invalid request", nil)
		return
	}

	message := services.MessageService.GetMessage(messageRequest)
	messageResponse := dto.MessageResponseBody(message)

	response.JsonResponse(rw, 200, "success", messageResponse)
}
