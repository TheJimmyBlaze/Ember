package authority

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thejimmyblaze/ember/common"
)

type Authority struct {
	config common.Config
	db     common.Database
}

func New(db common.Database, config common.Config) (*Authority, error) {

	log.Print("Configuring Authority...")

	auth := &Authority{
		config: config,
		db:     db,
	}

	return auth, nil
}

func (authority *Authority) GetConfig() common.Config {
	return authority.config
}

func (authority *Authority) GetDB() common.Database {
	return authority.db
}

func (authority *Authority) Shutdown() {
	authority.db.Close()
}
