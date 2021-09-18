package server

import (
	"go-starter/config"
	"go-starter/library/logger"
	"go-starter/storage/db"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var once sync.Once
var server Server

type Server struct {
	Router   *mux.Router
	Database db.MongoDB
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

func (s *Server) StartHttpServer() error {

	appPort := config.Get().AppPort
	srv := &http.Server{
		Handler:      s.Router,
		Addr:         ":" + appPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Log.Info("Server running on port: ", appPort)

	// TODO update with graceful shutdown
	err := srv.ListenAndServe()
	return err
}
