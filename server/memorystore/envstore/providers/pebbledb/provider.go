package pebbledb

import (
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
	log "github.com/sirupsen/logrus"
)

type provider struct {
	db *pebble.DB
}

func NewProvider() (*provider, error) {
	properties := pebble.Options{}

	properties.FS = vfs.NewMem()

	pebbledb, err := pebble.Open("", &properties)
	if err != nil {
		log.Debug("Failed to initialize pebbledb env store instance : ", err)
	}

	return &provider{
		db: pebbledb,
	}, err
}
