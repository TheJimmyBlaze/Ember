package authority

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thejimmyblaze/ember/config"
)

type Authority struct {
	Config *config.Config
	DB     *sql.DB
}

func New(config *config.Config) (*Authority, error) {

	db, err := configureDB(config)
	if err != nil {
		return nil, err
	}

	auth := &Authority{
		Config: config,
		DB:     db,
	}

	return auth, nil
}

func configureDB(config *config.Config) (*sql.DB, error) {

	log.Printf("Configuring DB: %s", config.DBFileName)

	db, err := sql.Open("sqlite3", config.DBFileName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	initTable := `
	create table authority (
		id integer not null,
		name text not null,

		constraint auth_id_pk primary key (id),
		constraint auth_name_unq unique (name)
	);`
	_, err = db.Exec(initTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}
