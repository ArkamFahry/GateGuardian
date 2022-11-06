package providers

import (
	"context"

	"github.com/ArkamFahry/GateGuardian/server/db/maindb/models"
)

type Provider interface {
	// User provider interfaces
	AddUser(ctx context.Context, user models.User) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUser()
	ListUser()
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserByID(ctx context.Context, id string) (models.User, error)
	UpdateUsers()

	// Session provider interface
	AddSession(ctx context.Context, session models.Session) error
}
