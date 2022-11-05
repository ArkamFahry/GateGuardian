package memdb

import (
	"github.com/hashicorp/go-memdb"
	"github.com/sirupsen/logrus"
)

type provider struct {
	db *memdb.MemDB
}

func NewMemDbProvider() (*provider, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"session": {
				Name: "session",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Key"},
					},
					"value": {
						Name:    "value",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Value"},
					},
				},
			},
		},
	}

	memdb, err := memdb.NewMemDB(schema)
	if err != nil {
		logrus.Fatal("memdb failed to create a new instance: ", err)
	}

	return &provider{
		db: memdb,
	}, err
}
