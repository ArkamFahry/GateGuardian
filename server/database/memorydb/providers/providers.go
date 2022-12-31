package providers

type Provider interface {
	// Updates the total env store
	UpdateEnvStore(store map[string]any) error

	// Gets the complete env store
	GetEnvStore() (map[string]any, error)

	// Updates a single variable in the env store
	UpdateEnvVariable(key string, value any) error

	// Gets the string env variables from env store
	GetStringStoreEnvVariable(key string) (string, error)
}
