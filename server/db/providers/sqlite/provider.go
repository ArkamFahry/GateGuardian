package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	_ "github.com/mattn/go-sqlite3"
)

type provider struct {
	db *sql.DB
}

func NewProvider() (*provider, error) {
	dbType, _ := envstore.Provider.GetEnv(constants.DB_TYPE)
	dbUrl, _ := envstore.Provider.GetEnv(constants.DB_URL)

	sqlite, err := sql.Open(dbType, dbUrl)

	if err != nil {
		return nil, err
	}

	createUsersCollectionQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (id TEXT PRIMARY KEY, email TEXT UNIQUE, password TEXT, given_name TEXT, family_name TEXT, middle_name TEXT, nick_name TEXT, gender TEXT, created_at INTEGER, updated_at INTEGER)`, models.Models.User)
	_, err = sqlite.Exec(createUsersCollectionQuery)
	if err != nil {
		return nil, err
	}

	return &provider{
		db: sqlite,
	}, nil
}
