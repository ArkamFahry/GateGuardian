package inmemory

import (
	"context"

	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/vfs"
	log "github.com/sirupsen/logrus"
)

type provider struct {
	ctx   context.Context
	store *pebble.DB
}

func NewMemoryStoreProvider(inmemoryPath string) (*provider, error) {
	var err error

	options := pebble.Options{}
	options.FS = vfs.NewMem()

	inmemory, err := pebble.Open(inmemoryPath, &options)
	if err != nil {
		log.Debug("error creating inmemory memorystore instance: ", err)
		return nil, err
	}
	ctx := context.Background()

	return &provider{
		ctx:   ctx,
		store: inmemory,
	}, nil
}
