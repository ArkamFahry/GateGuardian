package env

import (
	"errors"
	"gategaurdian/server/constants"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type RequiredEnv struct {
	DatabaseUrl string `json:"DATABASE_URL"`
	RedisUrl    string `json:"REDIS_URL"`
}

// RequiredEnvStore is a simple in-memory store for required envs
type RequiredEnvStore struct {
	mutex       sync.Mutex
	requiredEnv RequiredEnv
}

// GetRequiredEnv to gets a required env from the in-memory store
func (r *RequiredEnvStore) GetRequiredEnv() RequiredEnv {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.requiredEnv
}

// SetRequiredEnv to sets a required env into the in-memory store
func (r *RequiredEnvStore) SetRequiredEnv(requiredEnv RequiredEnv) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.requiredEnv = requiredEnv
}

// RequiredEnvStoreObj represents the in-memory required env store
var RequiredEnvStoreObj *RequiredEnvStore

func InitRequiredEnv() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Error("couldn't load env from .env : ", err)
	}

	dbUrl := viper.GetString(constants.EnvDatabaseUrl)
	redisUrl := viper.GetString(constants.EnvRedisUrl)

	if strings.TrimSpace(dbUrl) == "" {
		log.Debug("DATABASE_URL is not set")
		return errors.New("invalid database url. DATABASE_URL is required")
	}

	if strings.TrimSpace(redisUrl) == "" {
		log.Debug("REDIS_URL is not set")
		return errors.New("invalid redis url. REDIS_URL is required")
	}

	requiredEnv := RequiredEnv{
		DatabaseUrl: dbUrl,
		RedisUrl:    redisUrl,
	}

	RequiredEnvStoreObj = &RequiredEnvStore{
		requiredEnv: requiredEnv,
	}

	return nil
}
