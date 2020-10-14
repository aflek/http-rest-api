package model

import (
  "testing"
)

//TestUser - набор валидных данных, который используем при тесировании
func TestUser(t *testing.T) *User {
  t.Helper()

  return &User {
    Email:    "user@example.org",
    Password: "password",
  }
}
