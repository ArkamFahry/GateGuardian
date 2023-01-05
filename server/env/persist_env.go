package env

import (
	"context"
	"encoding/json"
	"gategaurdian/server/constants"
	"gategaurdian/server/crypto"
	"gategaurdian/server/database/db"
	"gategaurdian/server/database/db/models"
	"gategaurdian/server/database/memorystore"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GetEnvData returns the env data from database
func GetEnvData() (map[string]interface{}, error) {
	var result map[string]interface{}
	ctx := context.Background()
	env, err := db.Provider.GetEnv(ctx)
	// config not found in db
	if err != nil {
		log.Debug("Error while getting env data from db: ", err)
		return result, err
	}

	encryptionKey := env.Hash
	decryptedEncryptionKey, err := crypto.DecryptB64(encryptionKey)
	if err != nil {
		log.Debug("Error while decrypting encryption key: ", err)
		return result, err
	}

	err = memorystore.Provider.UpdateEnvVariable(constants.EnvEncryptionKey, decryptedEncryptionKey)
	if err != nil {
		log.Debug("Error while updating env encryption env variable : ", err)
		return result, err
	}

	b64DecryptedConfig, err := crypto.DecryptB64(env.Data)
	if err != nil {
		log.Debug("Error while decrypting env data from B64: ", err)
		return result, err
	}

	decryptedConfigs, err := crypto.DecryptAesByte([]byte(b64DecryptedConfig))
	if err != nil {
		log.Debug("Error while decrypting env data from AES: ", err)
		return result, err
	}

	err = json.Unmarshal(decryptedConfigs, &result)
	if err != nil {
		log.Debug("Error while unmarshalling env data: ", err)
		return result, err
	}

	return result, err
}

// PersistEnv persists the environment variables to the database
func PersistEnvData() error {
	ctx := context.Background()
	env, err := db.Provider.GetEnv(ctx)
	// config not found in db
	if err != nil || env.Data == "" {
		// AES encryption needs 32 bit key only, so we chop off last 4 characters from 36 bit uuid
		hash := uuid.New().String()[:36-4]
		err := memorystore.Provider.UpdateEnvVariable(constants.EnvEncryptionKey, hash)
		if err != nil {
			log.Debug("Error while updating encryption env variable: ", err)
			return err
		}
		encodedHash := crypto.EncryptB64(hash)

		res, err := memorystore.Provider.GetEnvStore()
		if err != nil {
			log.Debug("Error while getting env store: ", err)
			return err
		}

		encryptedConfig, err := crypto.EncryptEnvData(res)
		if err != nil {
			log.Debug("Error while encrypting env data: ", err)
			return err
		}

		env = models.Env{
			Hash: encodedHash,
			Data: encryptedConfig,
		}

		env, err = db.Provider.AddEnv(ctx, env)
		if err != nil {
			log.Debug("Error while persisting env data to db: ", err)
			return err
		}
	} else {
		// decrypt the config data from db
		// decryption can be done using the hash stored in db
		encryptionKey := env.Hash
		decryptedEncryptionKey, err := crypto.DecryptB64(encryptionKey)
		if err != nil {
			log.Debug("Error while decrypting encryption key: ", err)
			return err
		}

		err = memorystore.Provider.UpdateEnvVariable(constants.EnvEncryptionKey, decryptedEncryptionKey)
		if err != nil {
			log.Debug("Error while updating env encryption env variable : ", err)
			return err
		}

		b64DecryptedConfig, err := crypto.DecryptB64(env.Data)
		if err != nil {
			log.Debug("Error while decrypting env data from B64: ", err)
			return err
		}

		decryptedConfigs, err := crypto.DecryptAesByte([]byte(b64DecryptedConfig))
		if err != nil {
			log.Debug("Error while decrypting env data from AES: ", err)
			return err
		}

		// temp store variable
		storeData := map[string]interface{}{}

		err = json.Unmarshal(decryptedConfigs, &storeData)
		if err != nil {
			log.Debug("Error while un-marshalling env data: ", err)
			return err
		}

		// if env is changed via env file or OS env
		// give that higher preference and update db, but we don't recommend it

		hasChanged := false
		for key, value := range storeData {
			// don't override unexposed envs
			// check only for derivative keys
			// No need to check for ENCRYPTION_KEY which special key we use for encrypting config data
			// as we have removed it from json
			if key != constants.EnvEncryptionKey {
				viper.AutomaticEnv()
				envValue := strings.TrimSpace(viper.GetString(key))
				if envValue != "" {
					if value != nil && value.(string) != envValue {
						storeData[key] = envValue
						hasChanged = true
					}
				}
			}
		}

		err = memorystore.Provider.UpdateEnvStore(storeData)
		if err != nil {
			log.Debug("Error while updating env store: ", err)
			return err
		}

		if hasChanged {
			encryptedConfig, err := crypto.EncryptEnvData(storeData)
			if err != nil {
				log.Debug("Error while encrypting env data: ", err)
				return err
			}

			env.Data = encryptedConfig
			_, err = db.Provider.UpdateEnv(ctx, env)
			if err != nil {
				log.Debug("Failed to Update Config: ", err)
				return err
			}
		}
	}

	return nil
}
