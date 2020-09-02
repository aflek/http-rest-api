package apiserver

// APIServer struct
type APIServer struct {}

// New - функция инициализации API сервера
func New() *APIServer {
  return &APIServer{}
}

// Start - функция запуска API - сервера
func (s *APIServer) Start() error {
  return nil
}
