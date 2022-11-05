package memdb

import (
	"github.com/hashicorp/go-memdb"
	"github.com/sirupsen/logrus"
)

type provider struct {
	memorydb *memdb.MemDB
}

func NewMemDbProvider() (*provider, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"env": {
				Name: "env",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"data": {
						Name:    "data",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Data"},
					},
					"created_at": {
						Name:    "created_at",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "CreatedAt"},
					},
					"updated_at": {
						Name:    "updated_at",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "UpdatedAt"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		logrus.Fatal("memdb failed to create a new inmemory instance: ", err)
	}

	return &provider{
		memorydb: db,
	}, err
}
