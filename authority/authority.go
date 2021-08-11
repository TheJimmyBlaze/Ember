package authority

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thejimmyblaze/ember/config"
	"github.com/thejimmyblaze/ember/database"
)

type Authority struct {
	Config *config.Config
	DB     *database.EmberDB
}

func New(db *database.EmberDB, config *config.Config) (*Authority, error) {

	log.Print("Configuring Authority...")

	auth := &Authority{
		Config: config,
		DB:     db,
	}

	return auth, nil
}

func (authority *Authority) Shutdown() {
	authority.DB.Close()
}
