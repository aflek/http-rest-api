package sqlstore

import (
  "database/sql"

	"github.com/aflek/http-rest-api/internal/app/store"
	"github.com/aflek/http-rest-api/internal/app/model"
)

// UserRepository - структра репозитория User
type UserRepository struct {
  store *Store
}

// Create - метод добавления пользователя
func (r *UserRepository) Create(u *model.User) error {
  //Перед добавлением записи валидируем данные
  if err := u.Validate(); err != nil {
    return err
  }

  //Перед созданием записи выполняем обработку данных BeforeCreate: хэшируем пароль и т.п.
  if err := u.BeforeCreate(); err != nil {
    return err
  }

  //Добавляем данные
  return r.store.db.QueryRow(
    "INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
    u.Email,
    u.EncryptedPassword,
  ).Scan(&u.ID)
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

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
    if err == sql.ErrNoRows {
      return nil, store.ErrRecordNotFound
    }


    return nil, err
  }
  return u, nil
}
