package memorydb

import (
	"encoding/json"
	"gategaurdian/server/database/memorydb/providers"
	"gategaurdian/server/database/memorydb/providers/redis"

	log "github.com/sirupsen/logrus"
)

var Provider providers.Provider

// InitMemoryDb is used to initialize the caching database used by the application
func InitMemoryDb() error {
	var err error

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

	redisURL := requiredEnvs.RedisUrl
	log.Info("Initializing Redis memory store")
	Provider, err = redis.NewRedisProvider(redisURL)
	if err != nil {
		log.Fatal("Failed to initialize Redis driver: ")
		return err
	}

	// set default envs in redis
	Provider.UpdateEnvStore(defaultEnvs)

	return nil
}
