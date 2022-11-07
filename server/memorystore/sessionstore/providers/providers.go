package providers

type Provider interface {
	SetSession(key string, value string) (string, error)
	GetSession(key string) (string, error)
	DeleteSession(key string) error
}
