package main

import (
	"log"
	"github.com/aflek/http-rest-api/internal/app/apiserver"
)

func main() {
  //Создаем API-сервер
  s := apiserver.New()
  //Запускаем сервер
  if err:= s.Start; err != nil {
    log.Fatal(err)
  }
}
