package apiserver

import (
  "github.com/sirupsen/logrus"
)


// APIServer struct
type APIServer struct {
  config *Config
  logger *logrus.Logger
}

// New - функция инициализации API сервера
func New(config *Config) *APIServer {
  return &APIServer {
    config: config,
    logger: logrus.New(),
  }
}

// Start - функция запуска API - сервера
func (s *APIServer) Start() error {

  if err := s.configureLogger(); err != nil {
    return err
  }

  s.logger.Info("starting api server")

  return nil
}

//configureLogger - конфигурирование логирования
func (s *APIServer) configureLogger() error {
  level, err := logrus.ParseLevel(s.config.LogLevel)//парсим конфиг
  if err != nil {
    return err
  }

  s.logger.SetLevel(level) //задаем уровень логирования

  return nil

}
