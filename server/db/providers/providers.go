package providers

import (
	"github.com/ArkamFahry/GateGuardian/server/db/models"
)

type Provider interface {
	// User provider interfaces
	AddUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser()
	ListUser()
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id string) (models.User, error)
	UpdateUsers()

	// Session provider interface
	AddSession(session models.Session) error
}
