package inmemory

var (
	// State store prefix
	// stateStorePrefix = "gate_guardian_state:"

	// Env store prefix
	envStorePrefix = "gate_guardian_env"
)

// Gets the complete env store
func (c *provider) GetEnvStore() (map[string]any, error) {
	res := make(map[string]interface{})
	// data, err := c.store.HGetAll(c.ctx, envStorePrefix).Result()
	// if err != nil {
	// 	return nil, err
	// }
	// for key, value := range data {
	// 	res[key] = value
	// }

	return res, nil
}

// Updates the total env store
func (c *provider) UpdateEnvStore(store map[string]any) error {
	// for key, value := range store {
	// 	err := c.store.HSet(c.ctx, envStorePrefix, key, value).Err()
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

// Updates a single variable in the env store
func (c *provider) UpdateEnvVariable(key string, value any) error {
	// err := c.store.HSet(c.ctx, envStorePrefix, key, value).Err()
	// if err != nil {
	// 	log.Debug("Error saving redis token: ", err)
	// 	return err
	// }
	return nil
}

// Gets the string env variables from env store
func (c *provider) GetStringStoreEnvVariable(key string) (string, error) {
	// data, err := c.store.HGet(c.ctx, envStorePrefix, key).Result()
	// if err != nil {
	// 	return "", nil
	// }

	return "", nil
}
