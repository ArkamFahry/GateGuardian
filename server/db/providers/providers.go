package providers

import (
	"github.com/ArkamFahry/GateGuardian/server/db/models"
)

type Provider interface {
	// User provider interfaces
	AddUser(user models.User) (models.User, error)
	UpdateUser()
	DeleteUser()
	ListUser()
	GetUserByEmail(email string) (models.User, error)
	GetUserByID()
	UpdateUsers()
}
