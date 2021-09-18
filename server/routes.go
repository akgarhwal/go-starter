package server

import (
	"fmt"
	"go-starter/library/middleware"
	"go-starter/src/example/handlers"
	"go-starter/src/example/repositories"
	"go-starter/src/example/services"
	"net/http"
)

func (s *Server) routes() {

	s.Router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello from server")
	})

	// logging middleware which log total time taken to complete request
	s.Router.Use(middleware.LoggingMiddleware)

	// define extra routes
	exampleRepository := repositories.NewExampleRepostiory(s.Database)
	exampleService := services.NewExampleService(exampleRepository)
	exampleHandlers := handlers.NewHttpHandler(exampleService)
	// setup example routes
	s.Router.HandleFunc("/example", exampleHandlers.GetExample)
}
