package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Datastore interface {
	AllWines() ([]*Wine, error)
}

type DB struct {
	*sql.DB
}

func Connect(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
	//return db, nil
}

func Close(db *DB) error {
	err := db.Close()
	return err
}
