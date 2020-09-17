package apiserver

import (
  "github.com/aflek/http-rest-api/internal/app/store"
)

// Config ...
type Config struct {
  //Адрес на котором запускаем севрер
  BindAddr string `toml:"bind_addr"`
  //Уровень логирования
  LogLevel string `toml:"log_lavel"`
  //Подключение к БД
  Store *store.Config
}

//NewConfig - функция инициализации с дефолтными параметрами конфига
func NewConfig() *Config {
  return &Config {
    BindAddr: ":8080",
    LogLevel: "debug",
    Store:    store.NewConfig(),
  }
}
