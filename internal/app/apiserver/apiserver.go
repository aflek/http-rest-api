package apiserver

import (
	"io"
	"net/http"
  "github.com/sirupsen/logrus"
  "github.com/gorilla/mux"
  "github.com/aflek/http-rest-api/internal/app/store"
)


// APIServer struct
type APIServer struct {
  config *Config
  logger *logrus.Logger
  router *mux.Router
  store  *store.Store
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

  //подключение к БД
  if err := s.configureStore(); err != nil {
    return err
  }

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

//configureStore - хранилище
func (s *APIServer) configureStore() error {
  st := store.New(s.config.Store)
  if err := st.Open(); err != nil {
    return err
  }

  s.store = st

  return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {

  return func(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello")
  }

}
