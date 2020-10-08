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

  u, err := s.User().Create(&model.User{
    Email: "user&example.org",
  })
  assert.NoError(t, err)
  assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
  databaseURL := "host=localhost user=postgres password='123456' dbname=restapi_test sslmode=disable"
  email := "user@example.org"

  s, teardown := store.TestStore(t, databaseURL)
  defer teardown("users")

  //Тест на поиск несуществющего пользователя
  _, err := s.User().FindByEmail(email)
  assert.Error(t, err)

  //Создаем пользователя и тестируем его поиск
  s.User().Create(&model.User{
    Email: email,
  })

  u, err := s.User().FindByEmail(email)
  assert.NoError(t, err)
  assert.NotNil(t, u)

}
