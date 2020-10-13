package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
  "testing"
  "github.com/aflek/http-rest-api/internal/app/store/sqlstore"
  "github.com/aflek/http-rest-api/internal/app/model"

)

func TestUserRepository_Create(t *testing.T) {
  databaseURL := "host=localhost user=postgres password='123456' dbname=restapi_test sslmode=disable"

  db, teardown := sqlstore.TestDB(t, databaseURL)
  defer teardown("users")

  s := sqlstore.New(db)
  u := model.TestUser(t)
  assert.NoError(t, s.User().Create(u))
  assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
  databaseURL := "host=localhost user=postgres password='123456' dbname=restapi_test sslmode=disable"

  db, teardown := sqlstore.TestDB(t, databaseURL)
  defer teardown("users")

  s := sqlstore.New(db)
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
