package store

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

func GetNewStore() *Store {
    return &Store{}
}

type Store struct {
    db *sql.DB
}

func (s *Store) Open() error {
    db, err := sql.Open("postgres", "host=127.0.0.1 dbname=restapi_dev sslmode=disable user=postgres password=pass port=5432")
    if err != nil {
        fmt.Println(err)
        return err
    }

    if err := db.Ping(); err != nil {
        return err
    }
    // s.db = db
    s.db = db
    return nil
}

func (s *Store) InitDB() error {
    _, err := s.db.Query("CREATE TABLE IF NOT EXISTS  urls (uuid UUID primary key not null , originalURL varchar(255) not null , shortenedURL varchar(10) not null )")
    return err
}
