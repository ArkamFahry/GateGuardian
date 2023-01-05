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
	Port            string `json:"PORT"`
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

	port := viper.GetString(constants.EnvPort)
	dbType := viper.GetString(constants.EnvDatabaseType)
	dbUrl := viper.GetString(constants.EnvDatabaseUrl)
	memoryStoreType := viper.GetString(constants.EnvMemoryStoreType)
	memoryStoreUrl := viper.GetString(constants.EnvMemoryStoreUrl)

	// If the port is not set in env it will default to 8000
	if port == "" {
		port = "8000"
	}

	// If the dbType is not set in env it will default to sqlite
	if dbType == "" {
		dbType = constants.DbTypeSqlite
	}

	// Check if db url is set in env
	if strings.TrimSpace(dbUrl) == "" {
		// If dbType is sqlite and database url or in this case sqlite data path is not set in env it will default to data.db
		// If the dbType is other than sqlite it will throw an error
		if dbType == constants.DbTypeSqlite {
			dbUrl = "data.db"
		} else {
			log.Debug("DATABASE_URL is not set")
			return errors.New("invalid database url. DATABASE_URL is required")
		}
	}

	// If the memoryStoreType is not set it will default to inmemory
	if memoryStoreType == "" {
		memoryStoreType = constants.MemoryStoreTypeInmemory
	}

	// If the memoryStoreType is inmemory it will default the memoryStoreUrl to :memory:
	if memoryStoreType == constants.MemoryStoreTypeInmemory {
		memoryStoreUrl = ":memory:"
	}

	// Check if memoryStoreUrl is set in env
	if strings.TrimSpace(memoryStoreUrl) == "" {
		if memoryStoreType == constants.MemoryStoreTypeInmemory {
			memoryStoreUrl = ":memory:"
		} else {
			log.Debug("REDIS_URL is not set")
			return errors.New("invalid redis url. REDIS_URL is required")
		}
	}

	requiredEnv := RequiredEnv{
		Port:            port,
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
