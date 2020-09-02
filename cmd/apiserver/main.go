package main

import (
	"github.com/BurntSushi/toml"
	"flag"
  "log"
  "github.com/aflek/http-rest-api/internal/app/apiserver"
)

//Путь до конфига будем передавать флагом в командной строке
var (
  configPath string
)

func init() {
  flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
  flag.Parse()
  config := apiserver.NewConfig()
  _, err := toml.DecodeFile(configPath, config)
  if err != nil {
    log.Fatal(err)
  }

  //Создаем API-сервер
  s := apiserver.New(config)
  //Запускаем сервер
  if err:= s.Start(); err != nil {
    log.Fatal(err)
  }
}
