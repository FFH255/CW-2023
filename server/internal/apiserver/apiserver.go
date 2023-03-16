package apiserver

import (
	"net/http"

	"github.com/FFH255/CW-2023/internal/router"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
}

func New() *APIServer {
	return &APIServer{
		config: NewConfig(),
		logger: logrus.New(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	router := router.Configure()

	s.logger.Info("starting api server.")

	return http.ListenAndServe(s.config.Port, router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLever)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}
