package store_test

import (
	"github.com/stretchr/testify/assert"
  "testing"
  "github.com/aflek/http-rest-api/internal/app/store"
  "github.com/aflek/http-rest-api/internal/app/model"

)

func TestUserRepository_Create(t *testing.T) {
  databaseURL := "host=localhost user=postgres password='123456' dbname=restapi_test sslmode=disable"

  s, teardown := store.TestStore(t, databaseURL)
  defer teardown("users")

  u, err := s.User().Create(model.TestUser(t))
  assert.NoError(t, err)
  assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
  databaseURL := "host=localhost user=postgres password='123456' dbname=restapi_test sslmode=disable"

  s, teardown := store.TestStore(t, databaseURL)
  defer teardown("users")

  //Тест на поиск несуществющего пользователя
  email := "user@example.org"
  _, err := s.User().FindByEmail(email)
  assert.Error(t, err)

  //Создаем пользователя и тестируем его поиск
  u := model.TestUser(t)
  u.Email = email
  s.User().Create(u)

  u, err = s.User().FindByEmail(email)
  assert.NoError(t, err)
  assert.NotNil(t, u)

}
