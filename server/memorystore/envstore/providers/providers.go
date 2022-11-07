package providers

type Provider interface {
	SetEnv(key string, value string) (string, error)
	GetEnv(key string) (string, error)
	DeleteEnv(key string) error
}
