package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/ArkamFahry/GateGuardian/server/db/maindb/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

type provider struct {
	db *sql.DB
}

func NewSqliteProvider() (*provider, error) {
	sqlite, err := sql.Open("sqlite3", "../data.db")
	if err != nil {
		logrus.Fatal("Sqlite failed to create a new instance: ", err)
	}

	userCollectionQuery := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
			id TEXT PRIMARY KEY, 
			email TEXT UNIQUE, 
			email_verified_at INTEGER,
			password TEXT,
			sign_up_method TEXT,
			user_name TEXT,
			family_name TEXT,
			given_name TEXT,
			middle_name TEXT,
			nick_name TEXT,
			gender TEXT,
			birth_date TEXT,
			phone_number TEXT,
			phone_number_verified_at INTEGER,
			picture TEXT,
			roles TEXT,
			revoked_timestamp INTEGER,
			is_multi_factor_auth_enabled BOOL,
			updated_at INTEGER,
			created_at INTEGER,
			last_logged_in INTEGER
		)`, models.Collections.User)
	_, err = sqlite.Exec(userCollectionQuery)
	if err != nil {
		logrus.Fatal("Sqlite Failed to create table users: ", err)
	}

	sessionCollectionQuery := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
			id TEXT PRIMARY KEY,
			user_id TEXT UNIQUE,
			ip TEXT,
			created_at INTEGER,
			updated_at INTEGER
		)`, models.Collections.Session)
	_, err = sqlite.Exec(sessionCollectionQuery)
	if err != nil {
		logrus.Fatal("Sqlite Failed to create table session: ", err)
	}

	return &provider{
		db: sqlite,
	}, err
}
