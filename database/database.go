package database

import (
	"database/sql"
	"log"

	"github.com/thejimmyblaze/ember/config"
)

type EmberDB struct {
	SQL *sql.DB
}

func New(config *config.Config) (*EmberDB, error) {

	log.Printf("Preparing Ember DB: %s...", config.DBFileName)

	db, err := sql.Open("sqlite3", config.DBFileName)
	if err != nil {
		return nil, err
	}

	emberDB := EmberDB{
		SQL: db,
	}
	err = emberDB.initializeTables()
	if err != nil {
		return nil, err
	}

	return &emberDB, nil
}

func (db *EmberDB) Close() {
	db.SQL.Close()
}
