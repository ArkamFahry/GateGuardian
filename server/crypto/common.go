package crypto

import (
	"encoding/json"
	"gategaurdian/server/database/memorystore"
)

func EncryptEnvData(data map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	storeData, err := memorystore.Provider.GetEnvStore()
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(jsonBytes, &storeData)
	if err != nil {
		return "", err
	}

	configData, err := json.Marshal(storeData)
	if err != nil {
		return "", err
	}

	encryptedConfig, err := EncryptAesByte(configData)
	if err != nil {
		return "", err
	}

	return EncryptB64(string(encryptedConfig)), nil
}
