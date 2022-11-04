package genjidb

import (
	"fmt"

	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/genjidb/genji"
	"github.com/sirupsen/logrus"
)

type provider struct {
	db *genji.DB
}

func NewGenjiDbProvider() (*provider, error) {
	genjidb, err := genji.Open("../data")
	if err != nil {
		logrus.Fatal("GenjiDb failed to create a new inmemory instance: ", err)
	}

	// GenjiDb creates a new user collection
	userCollectionQuery := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
			id TEXT PRIMARY KEY, 
			email TEXT UNIQUE, 
			email_verified_at INTEGER,
			password TEXT,
			sign_up_method TEXT,
			user_name TEXT,
			first_name TEXT,
			middle_name TEXT,
			last_name TEXT,
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
	err = genjidb.Exec(userCollectionQuery)
	if err != nil {
		logrus.Fatal("GenjiDb Failed to create table users: ", err)
	}

	sessionCollectionQuery := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
			id TEXT PRIMARY KEY,
			user_id TEXT UNIQUE,
			ip TEXT,
			refresh_token TEXT,
			created_at INTEGER,
			updated_at INTEGER
		)`, models.Collections.Session)
	err = genjidb.Exec(sessionCollectionQuery)
	if err != nil {
		logrus.Fatal("GenjiDb Failed to create table session: ", err)
	}

	return &provider{
		db: genjidb,
	}, err
}
