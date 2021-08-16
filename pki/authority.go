package pki

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thejimmyblaze/ember/common"
	"github.com/thejimmyblaze/ember/config"
	"github.com/thejimmyblaze/ember/database"
)

type Authority struct {
	config common.Config
	db     common.Database
}

func CreateAuthority(configFileName string) (*Authority, error) {

	log.Print("Creating Authority...")

	config, err := config.New(configFileName)
	if err != nil {
		return nil, err
	}

	db, err := database.New(config)
	if err != nil {
		return nil, err
	}

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
