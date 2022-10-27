package providers

type Provider interface {
	// Env provider interfaces
	AddEnv()
	UpdateEnv()
	DeleteEnv()
	ListEnv()
	GetEnvByKey()
}
