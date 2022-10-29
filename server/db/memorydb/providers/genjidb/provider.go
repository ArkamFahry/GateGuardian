package genjidb

import (
	"github.com/genjidb/genji"
	"github.com/sirupsen/logrus"
)

type provider struct {
	memorydb *genji.DB
}

func NewGenjiDbProvider() (*provider, error) {
	genjidb, err := genji.Open(":memory:")
	if err != nil {
		logrus.Fatal("GenjiDb failed to create a new inmemory instance: ", err)
	}

	// GenjiDb creates a new env table
	err = genjidb.Exec(`
		CREATE TABLE env(
			id 			TEXT PRIMARY KEY, 
			data 	TEXT,
			created_at  INT,
			updated_at  INT 
		);
	`)
	if err != nil {
		logrus.Fatal("GenjiDb Failed to create table env: ", err)
	}

	return &provider{
		memorydb: genjidb,
	}, err
}
