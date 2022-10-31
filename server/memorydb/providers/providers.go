package providers

import "github.com/ArkamFahry/GateGuardian/server/memorydb/models"

type Provider interface {
	// Env provider interfaces
	AddEnv(key string, data string) (string, error)
	UpdateEnv(key string, data string) (string, error)
	DeleteEnv(key string) error
	ListEnv() ([]models.Env, error)
	GetEnvByKey(key string) (string, error)
}
