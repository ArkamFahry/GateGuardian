package inmemory

import (
	"fmt"

	"github.com/genjidb/genji/document"
	log "github.com/sirupsen/logrus"
)

var (
	// State store prefix
	// stateStorePrefix = "gate_guardian_state:"

	// Env store prefix
	envStorePrefix = "gate_guardian_env"
)

// Gets the complete env store
func (c *provider) GetEnvStore() (map[string]any, error) {
	res := make(map[string]interface{})
	query := fmt.Sprintf(`SELECT * FROM %s`, envStorePrefix)
	data, err := c.store.QueryDocument(query)
	if err != nil {
		return nil, err
	}

	err = document.MapScan(data, &res)
	if err != nil {
		log.Debug("Error mapping env data to interface failed : ", err)
		return nil, err
	}

	return res, nil
}

// Updates the total env store
func (c *provider) UpdateEnvStore(store map[string]any) error {
	query := fmt.Sprintf(`DELETE FROM %s`, envStorePrefix)
	err := c.store.Exec(query)
	if err != nil {
		return err
	}

	query = fmt.Sprintf(`INSERT INTO %s VALUES ?`, envStorePrefix)
	err = c.store.Exec(query, store)
	if err != nil {
		return err
	}

	return nil
}

// Updates a single variable in the env store
func (c *provider) UpdateEnvVariable(key string, value any) error {
	query := fmt.Sprintf(`UPDATE %s SET %s = ?`, envStorePrefix, key)
	err := c.store.Exec(query, value)
	if err != nil {
		log.Debug("Error saving to inmemory store : ", err)
		return err
	}
	return nil
}

// Gets the string env variables from env store
func (c *provider) GetStringStoreEnvVariable(key string) (string, error) {
	res := make(map[string]interface{})
	query := fmt.Sprintf(`SELECT %s FROM %s`, key, envStorePrefix)
	data, err := c.store.QueryDocument(query)
	if err != nil {
		return "", nil
	}

	err = document.MapScan(data, &res)
	if err != nil {
		log.Debug("Error mapping env data to interface failed : ", err)
		return "", err
	}

	env := res[key].(string)

	return env, nil
}
