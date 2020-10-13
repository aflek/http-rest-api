package sqlstore

import (
  "database/sql"
  "github.com/aflek/http-rest-api/internal/app/store"
  _ "github.com/lib/pq" //подключаем анонимно, что бы методы пакета Open, Close не импортировались в наш код,  т.к. у нас свои методы подключения и они не должны быть затерты методами из пакета
)

//Store ...
type Store struct {
  db             *sql.DB
  userRepository *UserRepository
}

//New - вспомогательный метод
// На вход передается БД
// Возвращает ссылку на хранилище
func New(db *sql.DB) *Store {
  return &Store{
    db: db,
  }
}

//User - метод обращения к таблице User для пользователей из "внешного мира" (не internal)
// пример вызова метода Ctrate: store.User().Create
func (s *Store) User() store.UserRepository {
  if s.userRepository != nil {
    return s.userRepository
  }
  //Если он не существует то инициализируем
  s.userRepository = &UserRepository{
    store: s,
  }

  return s.userRepository

}
