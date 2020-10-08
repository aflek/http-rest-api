package store

import (
  "database/sql"
  _ "github.com/lib/pq" //подключаем анонимно, что бы методы пакета Open, Close не импортировались в наш код,  т.к. у нас свои методы подключения и они не должны быть затерты методами из пакета
)

//Store ...
type Store struct {
  config         *Config
  db             *sql.DB
  userRepository *UserRepository
}

//New - вспомогательный метод
// На вход передается Config
// Возвращает ссылку на хранилище
func New(config *Config) *Store {
  return &Store{
    config: config,
  }
}

//Open - метод для инициализации хранилища - подключения к БД
func (s *Store) Open() error {
  db, err := sql.Open("postgres", s.config.DatabaseURL)
  if err != nil {
    return err
  }

  //Проверяем, что соединение установлено
  if err := db.Ping(); err != nil {
    return err
  }

  s.db = db

  return nil
}

//Close - метод отключения от БД
func (s *Store) Close() {
  s.db.Close()
}

//User - метод обращения к таблице User для пользователей из "внешного мира" (не internal)
// пример вызова метода Ctrate: store.User().Create
func (s *Store) User() *UserRepository {
  if s.userRepository != nil {
    return s.userRepository
  }
  //Если он не существует то инициализируем
  s.userRepository = &UserRepository{
    store: s,
  }

  return s.userRepository

}
