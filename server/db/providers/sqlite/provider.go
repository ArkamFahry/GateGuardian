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
	dbType := "sqlite3"
	dbUrl, _ := envstore.Provider.GetEnv(constants.DbUrl)

	sqlite, err := sql.Open(dbType, dbUrl)

	if err != nil {
		return nil, err
	}

	createUsersCollection := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id TEXT PRIMARY KEY, 
		email TEXT UNIQUE, 
		password TEXT, 
		given_name TEXT, 
		family_name TEXT, 
		middle_name TEXT, 
		nick_name TEXT, 
		gender TEXT, 
		birth_date TEXT,
		picture TEXT,
		created_at INTEGER, 
		updated_at INTEGER
	)`, models.Model.User)
	_, err = sqlite.Exec(createUsersCollection)
	if err != nil {
		return nil, err
	}

	createSessionsCollection := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id TEXT PRIMARY KEY, 
		user_id TEXT,
		user_agent TEXT,
		ip TEXT,
		created_at INTEGER, 
		updated_at INTEGER
	)`, models.Model.Session)
	_, err = sqlite.Exec(createSessionsCollection)
	if err != nil {
		return nil, err
	}

	createStateCollection := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id TEXT PRIMARY KEY, 
		user_id TEXT,
		code_challenge TEXT,
		auth_code TEXT,
		created_at INTEGER, 
		updated_at INTEGER
	)`, models.Model.State)
	_, err = sqlite.Exec(createStateCollection)
	if err != nil {
		return nil, err
	}

	return &provider{
		db: sqlite,
	}, nil
}
