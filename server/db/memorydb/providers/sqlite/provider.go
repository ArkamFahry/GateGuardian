package sqlite

import (
	"database/sql"

	"github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

type provider struct {
	memorydb *sql.DB
}

func NewSqliteProvider() (*provider, error) {
	sqlite, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		logrus.Fatal("SqlLite failed to create a new inmemory instance: ", err)
	}

	// SqlLite creates a new env table
	_, err = sqlite.Exec(`
		CREATE TABLE env(
			id 			TEXT PRIMARY KEY, 
			env_data 	TEXT 
		);
		INSERT INTO env (id, env_data) VALUES ('DB','mongodb');
	`)
	if err != nil {
		logrus.Fatal("SqlLite Failed to create table env: ", err)
	}

	return &provider{
		memorydb: sqlite,
	}, err
}
