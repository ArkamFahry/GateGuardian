package memorystore

import (
	"errors"
	"gategaurdian/server/constants"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type RequiredEnv struct {
	DatabaseType    string `json:"DATABASE_TYPE"`
	DatabaseUrl     string `json:"DATABASE_URL"`
	MemoryStoreType string `json:"MEMORYSTORE_TYPE"`
	MemoryStoreUrl  string `json:"MEMORYSTORE_URL"`
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

// Initialize the required env load the envs required for application startup
func InitRequiredEnv() error {
	// Viper is used for easy loading of config data
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Error("couldn't load env from .env : ", err)
		viper.AutomaticEnv()
	}

	dbType := viper.GetString(constants.EnvDatabaseType)
	dbUrl := viper.GetString(constants.EnvDatabaseUrl)
	memoryStoreType := viper.GetString(constants.EnvMemoryStoreType)
	memoryStoreUrl := viper.GetString(constants.EnvMemoryStoreUrl)

	if dbType == "" {
		dbType = constants.DbTypeSqlite
	}

	if strings.TrimSpace(dbUrl) == "" {
		if dbType == constants.DbTypeSqlite {
			dbUrl = "data.db"
		} else {
			log.Debug("DATABASE_URL is not set")
			return errors.New("invalid database url. DATABASE_URL is required")
		}
	}

	if memoryStoreType == "" {
		memoryStoreType = constants.MemoryStoreTypeInmemory
	}

	if strings.TrimSpace(memoryStoreUrl) == "" {
		if memoryStoreType == constants.MemoryStoreTypeInmemory {
			memoryStoreUrl = ""
		} else {
			log.Debug("REDIS_URL is not set")
			return errors.New("invalid redis url. REDIS_URL is required")
		}
	}

	requiredEnv := RequiredEnv{
		DatabaseType:    dbType,
		DatabaseUrl:     dbUrl,
		MemoryStoreType: memoryStoreType,
		MemoryStoreUrl:  memoryStoreUrl,
	}

	RequiredEnvStoreObj = &RequiredEnvStore{
		requiredEnv: requiredEnv,
	}

	return nil
}
