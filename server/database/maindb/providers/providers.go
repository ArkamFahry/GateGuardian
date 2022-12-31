package providers

type Provider interface {
	// AddEnv to save environment variable information in database
	AddEnv()

	// UpdateEnv to update environment variable information in database
	UpdateEnv()

	// GetEnv to get environment variable information from database
	GetEnv()
}
