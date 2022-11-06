package pebbledb

import (
	"context"

	"github.com/cockroachdb/pebble"
	log "github.com/sirupsen/logrus"
)

func (p *provider) SetSession(ctx context.Context, key string, value string) (string, error) {
	err := p.db.Set([]byte(key), []byte(value), pebble.Sync)
	if err != nil {
		log.Debug("pebbledb error can't set value")
	}

	return key, nil
}

func (p *provider) GetSession(ctx context.Context, key string) (string, error) {
	value, closer, err := p.db.Get([]byte(key))
	if err != nil {
		log.Debug("pebbledb error can't get value")
	}

	closer.Close()

	return string(value), nil
}

func (p *provider) UpdateSession(ctx context.Context, key string, value string) (string, error) {
	err := p.db.Set([]byte(key), []byte(value), pebble.Sync)
	if err != nil {
		log.Debug("pebbledb error can't update value")
	}

	return key, nil
}
func (p *provider) DeleteSession(ctx context.Context, key string) error {
	err := p.db.Delete([]byte(key), pebble.Sync)
	if err != nil {
		log.Debug("pebbledb error can't update value")
	}

	return nil
}
