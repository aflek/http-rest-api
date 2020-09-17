package store

//Config ...
type Config struct {
  DatabaseURL string `toml:"database_url"`
}

// NewConfig - функция, которая возвращает указатель на конфиг
func NewConfig() *Config {
  return &Config{}
}
