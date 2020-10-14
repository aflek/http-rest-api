package apiserver

import (
  "database/sql"
  "net/http"

  "github.com/aflek/http-rest-api/internal/app/store/sqlstore"
  "github.com/gorilla/sessions"
  _ "github.com/lib/pq"
)

// Start ...
func Start(config *Config) error {
  db, err := newDB(config.DatabaseURL)
  if err != nil {
    return err
  }

  defer db.Close()

  store := sqlstore.New(db)
  sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
  srv := newServer(store, sessionStore)

  return http.ListenAndServe(config.BindAddr, srv)
}

// newDB ...
func newDB(databaseURL string) (*sql.DB, error) {
  db, err := sql.Open("postgres", dbURL)
  if err != nil {
    return nil, err
  }

  defer db.Close()

  if err := db.Ping(); err != nil {
    return nil, err
  }

  return db, nil
}
