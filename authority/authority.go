package authority

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thejimmyblaze/ember/config"
)

type Authority struct {
	Config *config.Config
	DB     *sql.DB
}

func New(config *config.Config) (*Authority, error) {

	db, err := sql.Open("sqlite3", config.DBFileName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	auth := &Authority{
		Config: config,
		DB:     db,
	}

	return auth, nil
}
