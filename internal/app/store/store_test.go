package store_test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
  databaseURL := os.Getenv("DATABASE_URL")
  if databaseURL == "" {
    databaseURL = "host=localhost user=postgres password='123456' dbname=restapi_test sslmode=disable"
  }

  os.Exit(m.Run())
}

