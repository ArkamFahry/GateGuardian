package env

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func PersistEnv(envs Envs) error {
	var err error
	if envs.PORT == "" {
		envs.PORT = "8080"
	}
	envstore.Provider.SetEnv(constants.PORT, envs.PORT)

	if envs.DB_TYPE == "" {
		envs.DB_TYPE = "sqlite"
	}
	envstore.Provider.SetEnv(constants.DB_TYPE, envs.DB_TYPE)

	if envs.DB_URL == "" {
		envs.DB_URL = "../data.db"
	}
	envstore.Provider.SetEnv(constants.DB_URL, envs.DB_URL)

	if envs.JWT_SECRET == "" {
		envs.JWT_SECRET = (uuid.New().String() + uuid.New().String())
	}
	envstore.Provider.SetEnv(constants.JWT_SECRET, envs.JWT_SECRET)

	if envs.ACCESS_TOKEN_EXPIRY_TIME == "" {
		envs.ACCESS_TOKEN_EXPIRY_TIME = "15m"
	}
	envstore.Provider.SetEnv(constants.ACCESS_TOKEN_EXPIRY_TIME, envs.ACCESS_TOKEN_EXPIRY_TIME)

	clientID := envs.CLIENT_ID

	if clientID == "" {
		clientID = uuid.New().String()
	}
	envstore.Provider.SetEnv(constants.CLIENT_ID, clientID)

	algo := envs.JWT_TYPE

	if algo == "" {
		algo = "RS256"
	} else {
		if !crypto.IsHMACA(algo) && !crypto.IsRSA(algo) && !crypto.IsECDSA(algo) {
			log.Debug("Invalid JWT Algorithm")
		}
	}
	envstore.Provider.SetEnv(constants.JWT_TYPE, algo)

	private_key := envs.JWT_PRIVATE_KEY
	public_key := envs.JWT_PUBLIC_KEY

	if private_key == "" || public_key == "" {
		if crypto.IsRSA(algo) {
			_, private_key, public_key, _, err = crypto.NewRSAKey(algo, clientID)
			if err != nil {
				return err
			}

			envstore.Provider.SetEnv(constants.JWT_PRIVATE_KEY, private_key)
			envstore.Provider.SetEnv(constants.JWT_PUBLIC_KEY, public_key)
		} else if crypto.IsECDSA(algo) {
			_, private_key, public_key, _, err = crypto.NewECDSAKey(algo, clientID)
			if err != nil {
				return err
			}
			envstore.Provider.SetEnv(constants.JWT_PRIVATE_KEY, private_key)
			envstore.Provider.SetEnv(constants.JWT_PUBLIC_KEY, public_key)
		}
	} else {
		// parse keys to make sure they are valid
		if crypto.IsRSA(algo) {
			_, err = crypto.ParseRsaPrivateKeyFromPemStr(private_key)
			if err != nil {
				return err
			}

			_, err := crypto.ParseRsaPublicKeyFromPemStr(public_key)
			if err != nil {
				return err
			}

		} else if crypto.IsECDSA(algo) {
			_, err = crypto.ParseEcdsaPrivateKeyFromPemStr(private_key)
			if err != nil {
				return err
			}

			_, err := crypto.ParseEcdsaPublicKeyFromPemStr(public_key)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
