package database

import (
	"database/sql"
	"log"

	"github.com/thejimmyblaze/ember/config"
	shared "github.com/thejimmyblaze/ember/database/shared_interface"
)

type EmberDB struct {
	SQL *sql.DB
}

func New(config *config.Config) (*EmberDB, error) {

	log.Printf("Preparing Ember Database: %s...", config.DBFileName)

	db, err := sql.Open("sqlite3", config.DBFileName)
	if err != nil {
		return nil, err
	}

	emberDB := EmberDB{
		SQL: db,
	}
	err = emberDB.runMigrations()
	if err != nil {
		return nil, err
	}

	return &emberDB, nil
}

func (db *EmberDB) ExecuteTransaction(function shared.QueryFunction) error {

	transaction, err := db.Begin()
	if err != nil {
		return err
	}

	if err = function(db); err != nil {
		transactionErr := transaction.Rollback()
		if transactionErr != nil {
			log.Printf("Error rolling back transaction: %v", transactionErr)
		}
		return err
	}

	transaction.Commit()
	return nil
}

func (db *EmberDB) Begin() (*sql.Tx, error) {
	return db.SQL.Begin()
}

func (db *EmberDB) Execute(statement string, args ...interface{}) (sql.Result, error) {
	return db.SQL.Exec(statement, args...)
}

func (db *EmberDB) Query(statement string, args ...interface{}) (*sql.Rows, error) {
	return db.SQL.Query(statement, args...)
}

func (db *EmberDB) Close() {
	db.SQL.Close()
}
