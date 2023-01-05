package memorystore

import (
	"encoding/json"
	"gategaurdian/server/constants"
	"gategaurdian/server/database/memorystore/providers"
	"gategaurdian/server/database/memorystore/providers/inmemory"
	"gategaurdian/server/database/memorystore/providers/redis"

	log "github.com/sirupsen/logrus"
)

var Provider providers.Provider

// InitMemoryDb is used to initialize the caching database used by the application
func InitMemoryDb() error {
	var err error
	memoryStoreType := RequiredEnvStoreObj.GetRequiredEnv().MemoryStoreType

	defaultEnvs := map[string]any{}

	requiredEnvs := RequiredEnvStoreObj.GetRequiredEnv()
	requiredEnvMap := make(map[string]any)
	requiredEnvBytes, err := json.Marshal(requiredEnvs)
	if err != nil {
		log.Debug("Error while marshalling required envs: ", err)
		return err
	}
	err = json.Unmarshal(requiredEnvBytes, &requiredEnvMap)
	if err != nil {
		log.Debug("Error while unmarshalling required envs: ", err)
		return err
	}

	// merge default envs with required envs
	for key, val := range requiredEnvMap {
		defaultEnvs[key] = val
	}

	isInmemory := memoryStoreType != constants.MemoryStoreTypeRedis && memoryStoreType != constants.MemoryStoreTypeDragonFly && memoryStoreType != constants.MemoryStoreTypeKeyDb && memoryStoreType != constants.MemoryStoreTypeKvrocks
	isDedicated := memoryStoreType != constants.MemoryStoreTypeInmemory

	if isInmemory {
		memoryStoreUrl := requiredEnvs.MemoryStoreUrl
		log.Info("Initializing in memory store")
		Provider, err = inmemory.NewMemoryStoreProvider(memoryStoreUrl)
		if err != nil {
			log.Fatal("Failed to initialize in memory store: ")
			return err
		}
	}

	if isDedicated {
		memoryStoreUrl := requiredEnvs.MemoryStoreUrl
		log.Info("Initializing Redis memory store")
		Provider, err = redis.NewMemoryStoreProvider(memoryStoreUrl)
		if err != nil {
			log.Fatal("Failed to initialize Redis driver: ")
			return err
		}
	}

	// set default envs in redis
	Provider.UpdateEnvStore(defaultEnvs)

	return nil
}
