package providers

type Provider interface {
	// Env provider interfaces
	AddEnv(key string, data string) (string, error)
	UpdateEnv(key string, data string) (string, error)
	DeleteEnv(key string) error
	GetEnvByKey(key string) (string, error)
}
