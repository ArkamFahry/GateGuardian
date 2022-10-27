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
		logrus.Fatal(err)
	}
	defer genjidb.Close()

	return &provider{
		memorydb: genjidb,
	}, err
}
