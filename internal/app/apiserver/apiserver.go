package apiserver

import (
	"io"
	"net/http"
  "github.com/sirupsen/logrus"
  "github.com/gorilla/mux"
)


// APIServer struct
type APIServer struct {
  config *Config
  logger *logrus.Logger
  router *mux.Router
}

// New - функция инициализации API сервера
func New(config *Config) *APIServer {
  return &APIServer {
    config: config,
    logger: logrus.New(),
    router: mux.NewRouter(),
  }
}

// Start - функция запуска API - сервера
func (s *APIServer) Start() error {

  if err := s.configureLogger(); err != nil {
    return err
  }

  s.configureRouter()//подключем роутер

  s.logger.Info("starting api server")//логируем стар сервиса

  return http.ListenAndServe(s.config.BindAddr, s.router)
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

//configureRouter - роутер
func (s *APIServer) configureRouter() {
  s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {

  return func(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello")
  }

}
