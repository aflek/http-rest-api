package apiserver

// Config ...
type Config struct {
  //Адрес на котором запускаем севрер
  BindAddr string `toml:"bind_addr"`

}

//NewConfig - функция инициализации с дефолтными параметрами конфига
func NewConfig() *Config {
  return &Config {
    BindAddr: ":8080",
  }
}
