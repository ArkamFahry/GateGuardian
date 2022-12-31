package memorydb

import (
	"encoding/json"
	"gategaurdian/server/database/memorydb/providers"
	"gategaurdian/server/database/memorydb/providers/redis"
	"gategaurdian/server/env"

	log "github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitMemStore() error {
	var err error

	defaultEnvs := map[string]any{}

	requiredEnvs := env.RequiredEnvStoreObj.GetRequiredEnv()
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
		return err
	}

	// set default envs in redis
	Provider.UpdateEnvStore(defaultEnvs)

	return nil
}
