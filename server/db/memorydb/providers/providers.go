package providers

import "context"

type Provider interface {
	SetSession(ctx context.Context, key string, value string) (string, error)
	GetSession(ctx context.Context, key string) (string, error)
	UpdateSession(ctx context.Context, key string, value string) (string, error)
	DeleteSession(ctx context.Context, key string) error
}
