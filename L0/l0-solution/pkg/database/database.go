package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db     *sql.DB
	Config *ConfigDB
}

func NewDB(config *ConfigDB) *Store {
	return &Store{
		Config: config,
	}
}

func (s *Store) Connect() error {

	db, err := sql.Open("postgres", s.Config.databaseURL)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Disconnect() error {
	return s.db.Close()
}
