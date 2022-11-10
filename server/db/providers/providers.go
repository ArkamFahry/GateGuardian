package providers

import (
	"context"

	"github.com/ArkamFahry/GateGuardian/server/db/models"
)

type Provider interface {
	// User interfaces
	AddUser(ctx context.Context, user models.User) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, user models.User) error
	ListUsers(ctx context.Context, pagination models.Pagination) ([]models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserByID(ctx context.Context, id string) (models.User, error)
	UpdateUsers(ctx context.Context, data map[string]interface{}, ids []string) error

	// Session interface
	AddSession(ctx context.Context, session models.Session) error
}
