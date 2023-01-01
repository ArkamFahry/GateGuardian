package providers

import (
	"context"
	"gategaurdian/server/database/maindb/models"
)

type Provider interface {
	// AddEnv to save environment variable information in database
	AddEnv(c context.Context, env models.Env) (models.Env, error)

	// UpdateEnv to update environment variable information in database
	UpdateEnv(c context.Context, env models.Env) (models.Env, error)

	// GetEnv to get environment variable information from database
	GetEnv(c context.Context) (models.Env, error)
}
