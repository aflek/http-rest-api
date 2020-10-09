package model

import (
  "testing"
)

//TestUser - набор валидных данных, который используем при тесировании
func TestUser(t *testing.T) *User {
  return &User {
    Email:    "user@example.org",
    Password: "password",
  }
}
