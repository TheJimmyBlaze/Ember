package database

import (
	"database/sql"
	"log"

	"github.com/thejimmyblaze/ember/common"
)

type Database struct {
	SQL *sql.DB
}

func New(config common.Config) (*Database, error) {

	dbFileName := config.GetDBFileName()
	log.Printf("Preparing Ember Database: %s...", dbFileName)

	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		return nil, err
	}

	database := Database{
		SQL: db,
	}
	err = database.runMigrations()
	if err != nil {
		return nil, err
	}

	return &database, nil
}

func (db *Database) ExecuteTransaction(function common.QueryFunction) error {

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

func (db *Database) Begin() (*sql.Tx, error) {
	return db.SQL.Begin()
}

func (db *Database) Execute(statement string, args ...interface{}) (sql.Result, error) {
	return db.SQL.Exec(statement, args...)
}

func (db *Database) Query(statement string, args ...interface{}) (*sql.Rows, error) {
	return db.SQL.Query(statement, args...)
}

func (db *Database) Close() {
	db.SQL.Close()
}
