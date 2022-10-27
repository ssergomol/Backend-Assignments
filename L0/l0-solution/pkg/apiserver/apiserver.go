package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	router *mux.Router
	logger *logrus.Logger
	// database field
}

func NewServer(config *Config) *APIserver {
	server := &APIserver{
		config: config,
		router: mux.NewRouter(),
		logger: logrus.New(),
		// database
	}
	return server
}

func (server *APIserver) Start() error {
	if err := server.configureLogger(); err != nil {
		return err
	}
	server.configureRouter()

	server.logger.Info("starting streaming API server")
	return http.ListenAndServe(server.config.BindAddr, server.router)
}

func (server *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}

	logrus.SetLevel(level)
	return nil
}

func (server *APIserver) configureRouter() {
	server.RegisterHome()
}
