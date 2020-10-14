package sqlstore_test

import (
	"github.com/aflek/http-rest-api/internal/app/store"
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
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {
  databaseURL := "host=localhost user=postgres password='123456' dbname=restapi_test sslmode=disable"

  db, teardown := sqlstore.TestDB(t, databaseURL)
  defer teardown("users")

  s := sqlstore.New(db)
	u1 := model.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u1)
	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}
