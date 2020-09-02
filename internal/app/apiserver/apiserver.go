package apiserver

// APIServer struct
type APIServer struct {
  config *Config
}

// New - функция инициализации API сервера
func New(config *Config) *APIServer {
  return &APIServer{
    config: config,
  }
}

// Start - функция запуска API - сервера
func (s *APIServer) Start() error {
  return nil
}
