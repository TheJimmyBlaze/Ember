package common

import "database/sql"

type QueryFunction func(Database) error

type Database interface {
	ExecuteTransaction(QueryFunction) error
	Begin() (*sql.Tx, error)
	Execute(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	Close()
}
