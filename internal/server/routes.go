package server

import (
	"fmt"
	"go-starter/internal/middleware"
	"go-starter/src/message/handlers"
	"go-starter/src/message/repositories"
	"go-starter/src/message/services"
	"net/http"
)

func (s *Server) routes() {

	s.Router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello from server")
	})

	// logging middleware which log total time taken to complete request
	s.Router.Use(middleware.LoggingMiddleware)

	// define extra routes
	messageRepository := repositories.NewMessageRepostiory(s.Database)
	messageService := services.NewMessageService(messageRepository)
	messageHandlers := handlers.NewHttpHandler(messageService)
	// setup message routes
	s.Router.HandleFunc("/message", messageHandlers.GetMessage)
}
