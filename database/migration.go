package database

import (
	"log"

	"github.com/thejimmyblaze/ember/database/migration"
)

type Migration interface {
	GetVersion() int
	ApplyMigration() error
}

func (db *Database) runMigrations() error {

	migrations := [...]Migration{
		&migration.Migration001{
			DB: db,
		},
	}

	version, err := db.getSchemaVersion()
	if err != nil {
		return err
	}

	for _, migration := range migrations {
		if migration.GetVersion() > version {
			err := migration.ApplyMigration()
			if err != nil {
				return err
			}
		}
	}

	log.Printf("Database is up to date")

	return nil
}

func (db *Database) getSchemaVersion() (int, error) {

	statement := "select schema_version from version;"

	rows, err := db.Query(statement)
	if err != nil || !rows.Next() {
		if err.Error() == "no such table: version" {
			log.Print("No database detected, initializing new database...")
			return 0, nil
		}
		return 0, err
	}
	defer rows.Close()

	var version int
	err = rows.Scan(&version)
	log.Printf("Detected schema at version: %d", version)

	return version, err
}
