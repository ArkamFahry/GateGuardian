package sql

import (
	"context"
	"time"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb/models"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/google/uuid"
)

func (p *provider) AddUser(ctx context.Context, user models.User) (models.User, error) {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	if user.Roles == "" {
		defaultRole, _ := env.GetEnvByKey(constants.DefaultRoles)
		user.Roles = defaultRole
	}

	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()
	user.LastLoggedIn = time.Now().Unix()

	result := p.db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (p *provider) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	user.UpdatedAt = time.Now().Unix()

	result := p.db.Save(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (p *provider) DeleteUser() {

}

func (p *provider) ListUser() {

}

func (p *provider) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	result := p.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (p *provider) GetUserByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	result := p.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (p *provider) UpdateUsers() {

}
