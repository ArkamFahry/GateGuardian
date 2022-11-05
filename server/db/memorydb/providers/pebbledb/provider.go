package pebbledb

import (
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
	"github.com/sirupsen/logrus"
)

type provider struct {
	db *pebble.DB
}

func NewPebbleDbProvider() (*provider, error) {
	properties := pebble.Options{}
	properties.FS = vfs.NewMem()
	pebbledb, err := pebble.Open("", &properties)
	if err != nil {
		logrus.Fatal("pebbledb failed to create a new instance: ", err)
	}

	return &provider{
		db: pebbledb,
	}, err
}
