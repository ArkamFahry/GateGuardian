package inmemory

import (
	"context"
	"fmt"

	"github.com/genjidb/genji"
	log "github.com/sirupsen/logrus"
)

type provider struct {
	ctx   context.Context
	store *genji.DB
}

func NewMemoryStoreProvider(inmemoryPath string) (*provider, error) {
	var err error

	inmemory, err := genji.Open(inmemoryPath)
	if err != nil {
		log.Debug("error initializing embedded inmemory store : ", err)
		return nil, err
	}

	ctx := context.Background()

	inmemory = inmemory.WithContext(ctx)

	// Create env store query
	query := fmt.Sprintf(`CREATE TABLE %s`, envStorePrefix)
	err = inmemory.Exec(query)
	if err != nil {
		log.Debug("error creating inmemory env store : ", err)
		return nil, err
	}

	return &provider{
		ctx:   ctx,
		store: inmemory,
	}, nil
}
