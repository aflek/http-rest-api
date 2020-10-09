package store

import (
	"github.com/aflek/http-rest-api/internal/app/model"
)

// UserRepository - структра репозитория User
type UserRepository struct {
  store *Store
}

// Create - метод добавления пользователя
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
  //Перед добавлением записи валидируем данные
  if err := u.Validate(); err != nil {
    return nil, err
  }

  //Перед созданием записи выполняем обработку данных BeforeCreate: хэшируем пароль и т.п.
  if err := u.BeforeCreate(); err != nil {
    return nil, err
  }

  //Добавляем данные
  if err := r.store.db.QueryRow(
    "INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
    u.Email,
    u.EncryptedPassword,
  ).Scan(&u.ID); err != nil {
    return nil, err
  }
  //если нет ошибки, то возвращаем юзера
  return u, nil
}


// FindByEmail - метод поиска User по e-mail
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
  u := &model.User{}
  if err := r.store.db.QueryRow(
    "SELECT id, email, encrypted_password FROM users WHERE email = $1",
    email,
  ).Scan(
    &u.ID,
    &u.Email,
    &u.EncryptedPassword,
  ); err != nil {
    return nil, err
  }
  return u, nil
}
