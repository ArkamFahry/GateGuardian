package providers

import (
	"context"

	"github.com/ArkamFahry/GateGuardian/server/db/models"
)

type Provider interface {
	// User provider interfaces
	AddUser(ctx context.Context, user models.User) (models.User, error)
	UpdateUser()
	DeleteUser()
	ListUser()
	GetUserByEmail()
	GetUserByID()
	UpdateUsers()
}
