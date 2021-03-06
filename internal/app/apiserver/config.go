package apiserver


// Config ...
type Config struct {
  //Адрес на котором запускаем севрер
  BindAddr string `toml:"bind_addr"`
  //Уровень логирования
  LogLevel string `toml:"log_lavel"`
  //Подключение к БД
  DatabaseURL string `toml:"database_url"`
  SessionKey  string `toml:"session_key"`
}

//NewConfig - функция инициализации с дефолтными параметрами конфига
func NewConfig() *Config {
  return &Config {
    BindAddr: ":8080",
    LogLevel: "debug",
  }
}
