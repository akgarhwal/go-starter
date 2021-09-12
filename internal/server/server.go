package server

import (
	"go-starter/pkg/db"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var once sync.Once
var server Server

type Server struct {
	Router   *mux.Router
	Database db.MongoDB
	Log      *logrus.Logger
}

func NewServer() Server {
	once.Do(func() {
		server = Server{}
		server.Router = mux.NewRouter()
		server.routes()
	})
	return server
}

func (s *Server) SetDatabase(database db.MongoDB) {
	s.Database = database
}

func (s *Server) SetLogger(logger *logrus.Logger) {
	s.Log = logger
}

func (s *Server) StartHttpServer() error {

	appPort := "8080" //config.Get().AppPort
	srv := &http.Server{
		Handler:      s.Router,
		Addr:         ":" + appPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Server running on port: ", appPort)

	// TODO update with graceful shutdown
	err := srv.ListenAndServe()
	return err
}
